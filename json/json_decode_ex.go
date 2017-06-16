package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	//const jsonStream = `
	//	{"Name": "Ed", "Text": "Knock knock.", "Valid": true, "Index": 5, "IPAMArgs": {"ip": "192.168.1.1"}}
	//`
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock.", "Valid": true, "IPAMArgs": {"ip": "192.168.1.1"}}
	`
	type IPAMArgs struct {
		IP net.IP `json:"ip,omitempty"`
	}
	type Message struct {
		Name, Text string
		Valid      bool      `json:"Valid"`
        Index      *int      `json:"Index"`
		Args       *IPAMArgs `json:"-"`
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", m)
        if m.Index != nil {
		    fmt.Printf("%d\n", *m.Index)
        }
		//fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
