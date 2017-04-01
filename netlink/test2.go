package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	parent, err := netlink.LinkByName("eth10")
	if err != nil {
		fmt.Printf("Get host interface error: %v", err)
		return
	}
	fmt.Printf("%#v", parent)
}
