package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	links, _ := netlink.LinkList()

	for _, link := range links {
		fmt.Println(link)
	}
}
