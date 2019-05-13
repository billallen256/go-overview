package main

import (
	"log"
	"math/rand"
	"time"
)

type Packet struct {
	Name string
}

// START FUNC OMIT
func mergeChannels(tcpChan, udpChan <-chan *Packet, outputChan chan<- *Packet) {
	for {
		select {
		case packet := <-tcpChan:
			outputChan <- packet
		case packet := <-udpChan:
			outputChan <- packet
		}
	}
}

// END FUNC OMIT

func packetGenerator(name string, packetChan chan<- *Packet) {
	for {
		packetChan <- &Packet{name}
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func processPackets(packetChan <-chan *Packet) {
	for packet := range packetChan {
		log.Println(packet.Name)
	}
}

// START MAIN OMIT
func main() {
	tcpChan := make(chan *Packet, 5)
	udpChan := make(chan *Packet, 5)
	outputChan := make(chan *Packet, 5)
	go mergeChannels(tcpChan, udpChan, outputChan)
	go packetGenerator("tcp", tcpChan)
	go packetGenerator("udp", udpChan)
	go processPackets(outputChan)
	time.Sleep(10 * time.Second)
}

// END MAIN OMIT
