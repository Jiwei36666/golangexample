package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","172.17.42.200:36000")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", ":10001")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    i := 0
    for {
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
        n,err := Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
		n, addr, err := Conn.ReadFromUDP(buf[0:n])
		if err != nil {
			fmt.Println("rcv error:", err)
		}else{
			fmt.Println("Received ",string(buf[0:n]), " from ",addr)	
		}	
        time.Sleep(time.Second * 3)
    }
}

