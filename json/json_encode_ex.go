package main

import (
	"encoding/json"
	"fmt"
	"os"
    "net"
)

func main() {
    type IPAMArgs struct {
        IP net.IP `json:"ip,omitempty"`
    }

	type ColorGroup struct {
		ID     int `json:"id"`
		Name   string `json:"name"`
		Colors []string `json:"colors"`
        Valid  bool `json:"valid"`
        Args   *IPAMArgs `json:"-"`
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
        Valid: true,
        Args: &IPAMArgs{
                IP: net.ParseIP("192.168.1.1"),
            },
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
