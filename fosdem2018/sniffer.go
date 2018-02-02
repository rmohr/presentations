package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type DHCPAck struct {
	IP  net.IP
	MAC net.HardwareAddr
}

func main() {

	// PCAP START OMIT
	handle, _ := pcap.OpenLive("foo", 65536, true, pcap.BlockForever)
	defer handle.Close()

	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	// PCAP END OMIT

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
