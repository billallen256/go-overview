package pcap

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Reads a pcap file and calls func for each packet that was successfully read
func ReadPcap(name string, callback func(gopacket.Packet)) {
	pcapFile, err := pcap.OpenOffline(name)
	if err != nil {
		log.Println("Failed to open pcap:", err)
		return
	}

	defer pcapFile.Close()
	source := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())

	for packet := range source.Packets() {
		callback(packet)
	}
}
