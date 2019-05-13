package main

import (
	"bytes"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// START 1 OMIT
func ip4PacketsFromPcap(source *gopacket.PacketSource, packetChan chan<- gopacket.Packet) {
	defer close(packetChan)

	for packet := range source.Packets() {
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)

		if ip4Layer == nil {
			packetChan <- packet
		}
	}
}

// END 1 OMIT

// START 2 OMIT
func appLayerFromPackets(packetChan <-chan gopacket.Packet, appLayerChan chan<- gopacket.ApplicationLayer) {
	defer close(appLayerChan)

	for packet := range packetChan {
		appLayer := packet.ApplicationLayer()

		if appLayer != nil {
			appLayerChan <- appLayer
		}
	}
}

// END 2 OMIT

// START 3 OMIT
func scanAppLayer(appLayerChan <-chan gopacket.ApplicationLayer) {
	for appLayer := range appLayerChan {
		if bytes.Contains(appLayer.Payload(), []byte("ssh")) {
			log.Println("Found")
		} else {
			log.Println("Not Found")
		}
	}
}

// END 3 OMIT

// START MAIN OMIT
func main() {
	pcapFile, _ := pcap.OpenOffline("packets.pcap")
	defer pcapFile.Close()
	packetSource := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())

	packetChan := make(chan gopacket.Packet, 100)
	appLayerChan := make(chan gopacket.ApplicationLayer, 100)

	go ip4PacketsFromPcap(packetSource, packetChan)
	go appLayerFromPackets(packetChan, appLayerChan)

	scanAppLayer(appLayerChan)
}

// END MAIN OMIT
