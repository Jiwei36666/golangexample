package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	user   string
	pwd    string
	url    string
	cacert string
	token  string
)

func init() {
	flag.StringVar(&user, "u", "", "user name")
	flag.StringVar(&pwd, "p", "", "password")
	flag.StringVar(&url, "url", "https://google.com", "URL")
	flag.StringVar(&cacert, "cacert", "", "CA root path")
	flag.StringVar(&token, "token", "", "token value")

}

func main() {
	flag.Parse()
	tlsConf := &tls.Config{}
	if cacert != "" {
		certs := x509.NewCertPool()
		pemData, err := ioutil.ReadFile(cacert)
		if err != nil {
			fmt.Printf("read ca error: %v\n", err)
			return
		}
		certs.AppendCertsFromPEM(pemData)
		tlsConf.RootCAs = certs
	}

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConf}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if user != "" && pwd != "" {
		req.SetBasicAuth(user, pwd)
	}
	if token != "" {
		val := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", val)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Do error: %v\n", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("code:%d\n", res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("body:\n%v\n", string(body))
}
