package main

import (
	"net"
	"syscall"
	"fmt"

	"github.com/docker/libnetwork/ipvs"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

var (
	schedMethods = []string{
		ipvs.RoundRobin,
	}

	protocols = []string{
		"TCP",
	}

	fwdMethods = []uint32{
		ipvs.ConnectionFlagMasq,
	}

	fwdMethodStrings = []string{
		"Masq",
	}
)

func createDummyInterface() error {
	dummy := &netlink.Dummy{
		LinkAttrs: netlink.LinkAttrs{
			Name: "dummy",
		},
	}

	err := netlink.LinkAdd(dummy)
	if err != nil {
		return err
	}

	dummyLink, err := netlink.LinkByName("dummy")
	if err != nil {
		return err
	}

	ip, ipNet, err := net.ParseCIDR("10.1.1.1/24")
	if err != nil {
		return err
	}

	ipNet.IP = ip

	ipAddr := &netlink.Addr{IPNet: ipNet, Label: ""}
	err = netlink.AddrAdd(dummyLink, ipAddr)
	if err != nil {
		return err
	}
	return nil
}


func TestDestination() error{

	if err := createDummyInterface(); err != nil {
		return err
	}

	i, err := ipvs.New("")
	if err != nil {
		return err
	}

	for _, protocol := range protocols {

		s := ipvs.Service{
			AddressFamily: nl.FAMILY_V4,
			SchedName:     ipvs.RoundRobin,
		}

		switch protocol {
		case "FWM":
			s.FWMark = 1234
		case "TCP":
			s.Protocol = syscall.IPPROTO_TCP
			s.Port = 80
			s.Address = net.ParseIP("1.2.3.4")
			s.Netmask = 0xFFFFFFFF
		case "UDP":
			s.Protocol = syscall.IPPROTO_UDP
			s.Port = 53
			s.Address = net.ParseIP("2.3.4.5")
		}

		err := i.NewService(&s)
		if err != nil {
			return err
		}

		s.SchedName = ""
		for _, fwdMethod := range fwdMethods {
			d1 := ipvs.Destination{
				AddressFamily:   nl.FAMILY_V4,
				Address:         net.ParseIP("10.1.1.2"),
				Port:            5000,
				Weight:          1,
				ConnectionFlags: fwdMethod,
			}

			err := i.NewDestination(&s, &d1)
			if err != nil {
				return err
			}
			d2 := ipvs.Destination{
				AddressFamily:   nl.FAMILY_V4,
				Address:         net.ParseIP("10.1.1.3"),
				Port:            5000,
				Weight:          1,
				ConnectionFlags: fwdMethod,
			}

			err = i.NewDestination(&s, &d2)
			if err != nil {
				return err
			}
			d3 := ipvs.Destination{
				AddressFamily:   nl.FAMILY_V4,
				Address:         net.ParseIP("10.1.1.4"),
				Port:            5000,
				Weight:          1,
				ConnectionFlags: fwdMethod,
			}

			err = i.NewDestination(&s, &d3)
			if err != nil {
				return err
			}

			//err = i.DelDestination(&s, &d1)
			//err = i.DelDestination(&s, &d2)
			//err = i.DelDestination(&s, &d3)
		}
	}
	return nil
}

func main(){

	err := TestDestination()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} 
}
