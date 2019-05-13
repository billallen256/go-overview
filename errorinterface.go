package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// START OMIT
type PacketError struct {
	IpVersion   int
	PayloadSize int
}

func (pe PacketError) Error() string {
	return fmt.Sprintf("Error processing IPv%d, payload size %d bytes", pe.IpVersion, pe.PayloadSize)
}

func payloadSize(packet gopacket.Packet) (int, error) {
	// Assume bad things happened...
	return 0, PacketError{IpVersion: 4, PayloadSize: 20}
}

// END OMIT

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
