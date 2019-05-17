package main

import (
	"flag"
	"log"

	"github.com/gershwinlabs/go-overview/solutions/pkg/pcap"
)

func main() {
	name := flag.String("pcap", "packets.pcap", "PCAP file to process.")
	flag.Parse()

	pcap.ReadPacketMetas(*name, func(md *pcap.PacketMeta) {
		log.Println(md)
	})
}
