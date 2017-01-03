ckage main
import (
        "fmt"
        "net/http"
        "crypto/tls"
)

func main(){
        client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
        srv := "https://hub.oa.com/v1/users/"
        req, err := http.NewRequest("GET", srv, nil)
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }

        req.SetBasicAuth("xxxx", "xxxx")
        res, err := client.Do(req);
        if err != nil {
                fmt.Printf("Do error: %v\n", err)
                return
        }

        fmt.Printf("code:%d\n", res.StatusCode)
}
