package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"net"
	"os/exec"
)

func main() {
	iface := &netlink.Macvlan{
		LinkAttrs: netlink.LinkAttrs{Name: "macvlan0", ParentIndex: 2},
		Mode:      netlink.MACVLAN_MODE_BRIDGE,
	}

	netlink.LinkAdd(iface)

	mymacvlan, _ := netlink.LinkByName("macvlan0")

	// START OMIT
	dst := &net.IPNet{ // VM ip
		IP:   net.IPv4(192, 168, 200, 4),
		Mask: net.CIDRMask(32, 32),
	}

	gw := net.IPv4(192, 168, 200, 5) // via
	address := &netlink.Addr{
		IPNet: &net.IPNet{
			IP:   gw,
			Mask: net.CIDRMask(32, 32),
		},
	}
	netlink.AddrAdd(mymacvlan, address)
	netlink.LinkSetUp(mymacvlan)

	route := &netlink.Route{LinkIndex: mymacvlan.Attrs().Index, Dst: dst, Gw: gw}
	netlink.RouteAdd(route)

	out, _ := exec.Command("ip", "route").Output()
	fmt.Println(string(out))
	// END OMIT

	netlink.LinkDel(mymacvlan)
}
