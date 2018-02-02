package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {

	// START OMIT
	mymacvlan := &netlink.Macvlan{
		LinkAttrs: netlink.LinkAttrs{Name: "macvlan0", ParentIndex: 2},
		Mode:      netlink.MACVLAN_MODE_BRIDGE,
	}

	netlink.LinkAdd(mymacvlan)

	link, _ := netlink.LinkByName("macvlan0")
	fmt.Println(link)
	// END OMIT

	netlink.LinkDel(link)
}
