package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"os/exec"
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
	}
	foo, _ := netlink.LinkByName("foo")
	bar, _ := netlink.LinkByName("bar")

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
			LinkIndex: foo.Attrs().Index,
			Handle:    netlink.MakeHandle(0xffff, 0),
		},
	}
	netlink.QdiscAdd(qdisc)
	// QDISC END OMIT

	// TC START OMIT
	netlink.SetPromiscOn(bar)

	selectors := &netlink.TcU32Sel{
		Flags: netlink.TC_U32_TERMINAL,
		// match dhcp bar port 67
		Keys: []netlink.TcU32Key{{Off: 20, Val: 0x00000043, Mask: 0x0000ffff}},
	}
	filter := &netlink.U32{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: foo.Attrs().Index,
			Parent:    netlink.MakeHandle(0xffff, 0),
			Priority:  1,
			Protocol:  unix.ETH_P_IP,
		},
		Actions: []netlink.Action{netlink.NewMirredAction(bar.Attrs().Index)},
		Sel:     selectors,
	}
	netlink.FilterAdd(filter)
	out, _ := exec.Command("tc", "filter", "show", "dev", "foo", "ingress").Output()
	fmt.Println(string(out))
	// TC END OMIT

	netlink.LinkDel(foo)
	netlink.LinkDel(bar)
	netlink.LinkDel(mymacvlan)
}
