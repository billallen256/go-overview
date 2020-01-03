package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// START OMIT
const poolSize = 20  // tune to taste

func process(packetChan chan gopacket.Packet) {
	for packet := range packetChan {
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)

		if ip4Layer == nil {
			log.Println("Not an IPv4 packet")
			continue
		}
		// Good processing goes here...
	}
}

func main() {
	pcapFile, _ := pcap.OpenOffline("packets.pcap")
	defer pcapFile.Close()
	packetSource := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())
	packetChan := make(chan gopacket.Packet, poolSize)
	defer close(packetChan)

	for i := 0; i < poolSize; i++ { go process(packetChan) }
	for packet := range packetSource.Packets() { packetChan <- packet }
}

// END OMIT
