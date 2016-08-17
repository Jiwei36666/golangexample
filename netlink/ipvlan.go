package main

import (
	"fmt"
    "github.com/vishvananda/netlink"
)

func main() {
	hostIf, err := netlink.LinkByName("eth0")
	if err != nil {
		fmt.Printf("Get host interface error: %v", err)
		return
	}

    ipvlan := &netlink.IPVlan{
		LinkAttrs: netlink.LinkAttrs{
			Name: "ipvlan0",
			ParentIndex: hostIf.Attrs().Index,
		},
		Mode: netlink.IPVLAN_MODE_L3,
	}
    err = netlink.LinkAdd(ipvlan)
	if err != nil {
		fmt.Printf("add ipvlan interface error: %v", err)
	}
}
