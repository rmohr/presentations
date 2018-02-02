package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {

	// START OMIT
	mymacvlan := &netlink.Macvtap{
		Macvlan: netlink.Macvlan{
			LinkAttrs: netlink.LinkAttrs{Name: "foo", ParentIndex: 2},
			Mode:      netlink.MACVLAN_MODE_BRIDGE,
		},
	}

	netlink.LinkAdd(mymacvlan)

	link, _ := netlink.LinkByName("foo")
	fmt.Println(link)
	// END OMIT

	netlink.LinkDel(link)
}
