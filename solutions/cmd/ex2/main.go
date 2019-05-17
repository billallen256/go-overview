package main

import (
	"flag"
	"log"

	"github.com/gershwinlabs/go-overview/solutions/pkg/pcap"
	"github.com/google/gopacket"
)

func main() {
	name := flag.String("pcap", "packets.pcap", "PCAP file to process.")
	flag.Parse()

	pcap.ReadPcap(*name, func(p gopacket.Packet) {
		md := pcap.NewPacketMeta(p)
		if md != nil {
			log.Println(md)
		}
	})
}
