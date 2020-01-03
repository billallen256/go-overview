package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// START OMIT
func process(packet gopacket.Packet) {
	ip4Layer := packet.Layer(layers.LayerTypeIPv4)

	if ip4Layer == nil {
		log.Println("Not an IPv4 packet")
		return
	}
	// Good processing goes here...
}

func main() {
	pcapFile, _ := pcap.OpenOffline("packets.pcap")
	defer pcapFile.Close()
	packetSource := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())

	for packet := range packetSource.Packets() {
		go process(packet) // Caution: new goroutine per packet
	}
}

// END OMIT
