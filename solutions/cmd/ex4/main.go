package main

import (
	"flag"
	"log"
	"runtime"
	"sync"

	"github.com/gershwinlabs/go-overview/solutions/pkg/pcap"
	"github.com/google/gopacket"
)

func main() {
	name := flag.String("pcap", "packets.pcap", "PCAP file to process.")
	parallel := flag.Int("parallel", runtime.NumCPU(), "Number of goroutines to use for computation.")
	flag.Parse()

	in := make(chan gopacket.Packet)
	results := make(chan map[string]int, *parallel)
	done := &sync.WaitGroup{}
	done.Add(*parallel)

	// Count in parallel
	for i := 0; i < *parallel; i++ {
		go func() {
			counts := make(map[string]int)

			for p := range in {
				md, err := pcap.NewPacketMeta(p)

				if err == nil {
					counts[md.Protocol]++
				}
			}

			results <- counts
			done.Done()
		}()
	}

	// Read pcap and send to counting routines
	pcap.ReadPackets(*name, func(p gopacket.Packet) {
		in <- p
	})
	close(in)

	// Collect results
	done.Wait()
	counts := make(map[string]int)

	for i := 0; i < *parallel; i++ {
		c := <-results

		for proto, count := range c {
			counts[proto] += count
		}
	}

	// Print results
	for proto, count := range counts {
		log.Println(proto, "->", count)
	}
}
