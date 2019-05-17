package main

import (
	"flag"
	"log"

	"github.com/gershwinlabs/go-overview/solutions/pkg/pcap"
)

func main() {
	name := flag.String("pcap", "packets.pcap", "PCAP file to process.")
	flag.Parse()

	counts := make(map[string]int)
	pcap.ReadPacketMetas(*name, func(md *pcap.PacketMeta) {
		counts[md.Protocol]++
	})
	for proto, count := range counts {
		log.Println(proto, "->", count)
	}
}
