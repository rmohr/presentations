package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {

	// START OMIT
	mymacvtap := &netlink.Macvlan{
		LinkAttrs: LinkAttrs{Name: "foo", ParentIndex: 2},
		Mode:      netlink.MACVLAN_MODE_BRIDGE,
	}
	netlink.LinkAdd(mymacvtap)
	link, _ := netlink.LinkByName("foo")
	fmt.Println(link)
	// END OMIT
}
