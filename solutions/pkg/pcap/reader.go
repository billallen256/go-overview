package pcap

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// ReadPackets reads a pcap file and calls func for each packet that was successfully read.
func ReadPackets(name string, callback func(gopacket.Packet)) {
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

// ReadPacketMetas reads a pcap file and calls func for each packet meta that was sucessfully read.
func ReadPacketMetas(name string, callback func(*PacketMeta)) {
	ReadPackets(name, func(p gopacket.Packet) {
		md := NewPacketMeta(p)
		if md != nil {
			callback(md)
		}
	})
}
