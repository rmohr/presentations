package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"os/exec"
)

func main() {

	// BRIDGE START OMIT
	mybridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{Name: "sniffer0"},
	}

	netlink.LinkAdd(mybridge)

	bridge, _ := netlink.LinkByName("sniffer0")
	netlink.LinkSetUp(bridge)
	netlink.SetPromiscOn(bridge)
	bridge, _ = netlink.LinkByName("sniffer0")
	out, _ := exec.Command("ip", "a").Output()
	fmt.Println(string(out))
	// BRIDGE END OMIT

	bridge, _ = netlink.LinkByName("sniffer0")
	netlink.LinkDel(bridge)
}
