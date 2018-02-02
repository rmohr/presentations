package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {

	// BRIDGE START OMIT
	for _, name := range []string{"foo", "bar"} {
		mybridge := &netlink.Bridge{
			LinkAttrs: netlink.LinkAttrs{Name: name},
		}

		netlink.LinkAdd(mybridge)

		bridge, _ := netlink.LinkByName(name)
		netlink.LinkSetUp(bridge)
		bridge, _ = netlink.LinkByName(name)
		fmt.Println(bridge)
	}
	// BRIDGE END OMIT

	bridge, _ := netlink.LinkByName("foo")
	netlink.LinkDel(bridge)

	bridge, _ = netlink.LinkByName("bar")
	netlink.LinkDel(bridge)
}
