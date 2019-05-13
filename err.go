package main

import (
	"fmt"
	"log"
)

// START MAIN OMIT
import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	pcapFile, _ := pcap.OpenOffline("packets.pcap")
	defer pcapFile.Close()
	packetSource := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())

	for packet := range packetSource.Packets() {
		s, err := payloadSize(packet)

		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
			continue
		}

		log.Printf("Payload Size = %d bytes", s)
	}
}

// END MAIN OMIT

// START FUNC OMIT
func payloadSize(packet gopacket.Packet) (int, error) {
	ip4Layer := packet.Layer(layers.LayerTypeIPv4)

	if ip4Layer == nil {
		return 0, fmt.Errorf("Not an IPv4 packet")
	}

	appLayer := packets.ApplicationLayer()

	if appLayer == nil {
		return 0, fmt.Errorf("Application Layer not available")
	}

	return len(appLayer.Payload()), nil
}

// END FUNC OMIT
