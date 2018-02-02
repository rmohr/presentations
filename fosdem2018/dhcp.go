package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"os/exec"
)

func main() {

	// BRIDGE START OMIT
	mybridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{Name: "foo"},
	}

	netlink.LinkAdd(mybridge)

	bridge, _ := netlink.LinkByName("foo")
	netlink.LinkSetUp(bridge)
	bridge, _ = netlink.LinkByName("foo")
	foo, _ := netlink.LinkByName("foo")

	mymacvlan := &netlink.Macvlan{
		LinkAttrs: netlink.LinkAttrs{Name: "test", ParentIndex: foo.Attrs().Index},
		Mode:      netlink.MACVLAN_MODE_BRIDGE,
	}
	netlink.LinkAdd(mymacvlan)
	netlink.LinkSetUp(mymacvlan)
	// BRIDGE END OMIT

	// QDISC START OMIT

	qdisc := &netlink.Ingress{
		QdiscAttrs: netlink.QdiscAttrs{
			Parent:    netlink.HANDLE_INGRESS,
			LinkIndex: 2,
			Handle:    netlink.MakeHandle(0xffff, 0),
		},
	}
	netlink.QdiscAdd(qdisc)
	// QDISC END OMIT

	// TC START OMIT
	netlink.SetPromiscOn(foo)

	selectors := &netlink.TcU32Sel{
		Flags: netlink.TC_U32_TERMINAL,
		// match dhcp  port 68 on device with index 2
		Keys: []netlink.TcU32Key{{Off: 20, Val: 0x00000044, Mask: 0x0000ffff}},
	}
	filter := &netlink.U32{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: 2,
			Parent:    netlink.MakeHandle(0xffff, 0),
			Priority:  1,
			Protocol:  unix.ETH_P_IP,
		},
		Actions: []netlink.Action{netlink.NewMirredAction(foo.Attrs().Index)},
		Sel:     selectors,
	}
	netlink.FilterAdd(filter)
	out, _ := exec.Command("tc", "filter", "show", "dev", "enp0s25", "ingress").Output()
	fmt.Println(string(out))
	// TC END OMIT

	exec.Command("tc", "filter", "del", "dev", "enp0s25", "parent", "ffff:").Output()
	netlink.LinkDel(foo)
	netlink.LinkDel(mymacvlan)
}
