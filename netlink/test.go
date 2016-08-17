package main

import (
	"fmt"
    "github.com/vishvananda/netlink"
)

func main() {
    la := netlink.NewLinkAttrs()
    la.Name = "foo"
    mybridge := &netlink.Bridge{la}
    err := netlink.LinkAdd(mybridge)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
    eth1, err := netlink.LinkByName("dummy")
    netlink.LinkSetMaster(eth1, mybridge)
}
