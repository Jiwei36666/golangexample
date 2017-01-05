//echo_epoll_server.c
#include "helper.h"
#include <sys/epoll.h>
#include <fcntl.h>
#include <sys/un.h>

#include <systemd/sd-daemon.h>

#define EVENT_ARR_SIZE 20
#define EPOLL_SIZE     20

void setnonblocking(
    int sockfd
);

int
main(int argc, char **argv)
{
    int        i,  listenfd, connfd, sockfd, epfd;
    ssize_t        n;
    char            buf[MAXLINE];
    socklen_t        clilen;
    struct sockaddr_in    cliaddr, servaddr;
    struct epoll_event ev, evs[EVENT_ARR_SIZE];
    int   nfds;

    n = sd_listen_fds(0);
    if (n > 1) {
        fprintf(stderr, "Too many file descriptors received.\n");
        exit(1);
    } else if (n == 1)
        listenfd = SD_LISTEN_FDS_START + 0;
    else {
        union {
                struct sockaddr sa;
                struct sockaddr_un un;
        } sa;

        listenfd = socket(AF_UNIX, SOCK_STREAM, 0);
        if (listenfd < 0) {
                fprintf(stderr, "socket(): %m\n");
                exit(1);
        }
    	setnonblocking(listenfd);

        memset(&sa, 0, sizeof(sa));
        sa.un.sun_family = AF_UNIX;
        strncpy(sa.un.sun_path, "/run/foobar.sk", sizeof(sa.un.sun_path));

        if (bind(listenfd, &sa.sa, sizeof(sa)) < 0) {
                fprintf(stderr, "bind(): %m\n");
                exit(1);
        }

        if (listen(listenfd, SOMAXCONN) < 0) {
                fprintf(stderr, "listen(): %m\n");
                exit(1);
        }
    }
    
    epfd = epoll_create(EPOLL_SIZE);
    ev.data.fd = listenfd;
    ev.events = EPOLLIN | EPOLLET;
    if(epoll_ctl(epfd, EPOLL_CTL_ADD, listenfd, &ev) < 0)
        err_sys("epoll_ctl listenfd error!\n");
    
    printf("server is listening....\n");

    for ( ; ; ) {
        if((nfds = epoll_wait(epfd, evs, EVENT_ARR_SIZE, -1)) < 0)
            err_sys("epoll_wait error!\n");

        for(i = 0; i < nfds; i++)
        {
                if(evs[i].data.fd == listenfd)
                {
                    clilen = sizeof(cliaddr);
                    connfd = accept(listenfd, (struct sockaddr*) &cliaddr, &clilen);
                    if(connfd < 0)
                        continue;
                        
                    setnonblocking(connfd);
                    ev.data.fd = connfd;
                    ev.events = EPOLLIN | EPOLLET;
                    if (epoll_ctl(epfd, EPOLL_CTL_ADD, connfd, &ev) < 0)
                        err_sys("epoll_ctl connfd error!\n");            
                }
                else if(evs[i].events & EPOLLIN)
                {
                    sockfd = evs[i].data.fd;
                    if (sockfd < 0)
                        continue;
                    if ( (n = read(sockfd, buf, MAXLINE)) == 0) {
                        epoll_ctl(epfd, EPOLL_CTL_DEL, sockfd, &ev);
                        close(sockfd);
                        evs[i].data.fd = -1;
                    } 
                    else if(n < 0)
                        err_sys("read socket error!\n");
                    else
                    {
                        printf("write %d bytes\n", n);
                        write(sockfd, buf, n);
                    }
                }
                else
                    printf("other event!\n");
        }
    }
    return 0;
}


void setnonblocking(
    int sockfd
)
{
    int flag;
    
    flag = fcntl(sockfd, F_GETFL);
    if(flag < 0)
            err_sys("fcnt(F_GETFL) error!\n");
    flag |= O_NONBLOCK;
    if(fcntl(sockfd, F_SETFL, flag) < 0)
        err_sys("fcon(F_SETFL) error!\n");
}

