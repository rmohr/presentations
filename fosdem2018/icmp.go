package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {

	// BRIDGE START OMIT
	mybridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{Name: "foo"},
	}

	netlink.LinkAdd(mybridge)

	bridge, _ := netlink.LinkByName("foo")
	fmt.Println(bridge)
	netlink.LinkSetUp(bridge)
	bridge, _ = netlink.LinkByName("foo")
	fmt.Println(bridge)
	// BRIDGE END OMIT

	netlink.LinkDel(bridge)
}
