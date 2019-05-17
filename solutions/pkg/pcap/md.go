package pcap

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// PacketMeta represents metadata about a packet.
type PacketMeta struct {
	Protocol    string
	SrcIp       net.IP
	DstIp       net.IP
	PayloadSize int
}

// String formats the metadata for output.
func (p *PacketMeta) String() string {
	return fmt.Sprint(
		"{Protocol:", p.Protocol,
		", SrcIp:", p.SrcIp,
		", DstIp:", p.DstIp,
		", PayloadSize:", p.PayloadSize,
		"}")
}

// NewPacketMeta creates a new PacketMeta from a packet.
// Returns error if the packet was not valid.
func NewPacketMeta(packet gopacket.Packet) (*PacketMeta, error) {
	ip4Layer := packet.Layer(layers.LayerTypeIPv4)
	payload := packet.ApplicationLayer()

	if ip4Layer == nil || payload == nil {
		return nil, fmt.Errorf("Invalid packet")
	}

	ip4 := ip4Layer.(*layers.IPv4)

	md := PacketMeta{
		Protocol:    ip4.NextLayerType().String(),
		SrcIp:       ip4.SrcIP,
		DstIp:       ip4.DstIP,
		PayloadSize: len(payload.Payload()),
	}

	return &md, nil
}
