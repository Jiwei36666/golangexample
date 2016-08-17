package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	parent, err := netlink.LinkByName("eth0")
	if err != nil {
		fmt.Printf("Get host interface error: %v", err)
		return
	}

	vxlan := &netlink.Vxlan{
		LinkAttrs:    netlink.LinkAttrs{Name: "vxlan1"},
		VxlanId:      258,
		VtepDevIndex: parent.Attrs().Index,
		Port:         8472,
	}
	err = netlink.LinkAdd(vxlan)
	if err != nil {
		fmt.Printf("add vxlan interface error: %v", err)
	}
}
