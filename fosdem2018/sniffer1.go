package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw"
	"github.com/vishvananda/netlink"
)

type DHCPAck struct {
	IP  net.IP
	MAC net.HardwareAddr
}

func main() {

	bridge, _ := netlink.LinkByName("snifferX")
	netlink.LinkDel(bridge)

	mybridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{Name: "snifferX"},
	}

	netlink.LinkAdd(mybridge)
	bridge, _ = netlink.LinkByName("snifferX")
	netlink.SetPromiscOn(bridge)

	ip := net.IPv4(192, 168, 200, 99) // via
	address := &netlink.Addr{
		IPNet: &net.IPNet{
			IP:   ip,
			Mask: net.CIDRMask(32, 32),
		},
	}
	netlink.AddrAdd(bridge, address)
	netlink.LinkSetUp(bridge)

	handle, err := OpenLive("snifferX")
	if err != nil {
		return
	}

	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	// PCAP END OMIT

	// PROCESS START OMIT
	for {
		var packet gopacket.Packet
		var open bool
		select {
		case packet, open = <-in:

			if !open {
				return
			}

			// Check if we deal with a dhcp server response
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer == nil {
				continue
			}
			udp := udpLayer.(*layers.UDP)
			udp.SetNetworkLayerForChecksum(packet.NetworkLayer())

			if udp.DstPort != 68 {
				// Not for a dhcp client
				continue
			}

			// PROCESS END OMIT

			// Check if we are dealing with a DHCP response
			dhcpLayer := packet.Layer(layers.LayerTypeDHCPv4)
			if dhcpLayer == nil {
				continue
			}

			dhcp := dhcpLayer.(*layers.DHCPv4)

			// If we have a DHCP ack, inform us about the discovered details
			if getOptionType(dhcp.Options) == layers.DHCPMsgTypeAck {
				fmt.Println("%v,%v", dhcp.YourClientIP, dhcp.ClientHWAddr)
			}

		}
	}
}

func getOptionType(options layers.DHCPOptions) layers.DHCPMsgType {
	for _, o := range options {
		if o.Type == layers.DHCPOptMessageType {
			return layers.DHCPMsgType(o.Data[0])
		}
	}

	return layers.DHCPMsgTypeUnspecified
}

// RAW START OMIT
func OpenLive(ifname string) (*Handle, error) {
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		return nil, err
	}
	c, err := raw.ListenPacket(iface, 0xcccc, nil) // HL
	if err != nil {
		return nil, err
	}
	return &Handle{C: c, MTU: iface.MTU}, nil
}
func (p *Handle) ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error) {
	data = make([]byte, p.MTU)
	n, _, err := p.C.ReadFrom(data) // HL
	ci = gopacket.CaptureInfo{CaptureLength: n}
	return data, ci, err
}

// RAW END OMIT

// RAW WRITE START OMIT
func (p *Handle) WritePacketData(data []byte) (err error) {
	addr := &raw.Addr{
		HardwareAddr: ethernet.Broadcast,
	}
	_, err = p.C.WriteTo(data, addr) // HL
	return err
}

// RAW WRITE END OMIT

type Handle struct {
	C   net.PacketConn
	MTU int
}
