package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	link, _ := netlink.LinkByName("lo")
	addrs, _ := netlink.AddrList(link, netlink.FAMILY_V4)

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
