package main
import (
        "fmt"
		"os"
        "net/http"
        "crypto/tls"
)

func main(){
        client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
        srv := os.Args[1]
        req, err := http.NewRequest("GET", srv, nil)
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }

		if len(os.Args) >= 4 {
        	req.SetBasicAuth(os.Args[2], os.Args[3])
		}
        res, err := client.Do(req);
        if err != nil {
                fmt.Printf("Do error: %v\n", err)
                return
        }

        fmt.Printf("code:%d\n", res.StatusCode)
}
