package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	srvAddr string
	msgSize int
)

func init() {
	flag.IntVar(&msgSize, "size", 1472, "UDP message size per packet")
	flag.StringVar(&srvAddr, "srv-addr", "127.0.0.1:40000", "Server address: ip and port")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	flag.Parse()

	ServerAddr, err := net.ResolveUDPAddr("udp", srvAddr)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()

	//buf := make([]byte, msgSize)
	i := 0
	for {
		s := fmt.Sprintf("test %d", i)
		i += 1

		buf := []byte(s)
		n, err := Conn.Write(buf)
		if err != nil {
			fmt.Println("write data error:", err)
			return
		} else {
			fmt.Printf("write %d bytes\n", n)
		}
		n, addr, err := Conn.ReadFromUDP(buf[0:n])
		if err != nil {
			fmt.Println("rcv error:", err)
		} else {
			fmt.Printf("received '%s' from %v\n", string(buf[0:n]), addr)
			//fmt.Println("received ", n, " bytes from ", addr)
		}

		time.Sleep(time.Second)
	}
}
