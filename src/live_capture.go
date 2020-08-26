package src

import (
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
	size         int = 0
	upBytes          = 0
	downBytes        = 0
	packetUp     gopacket.Packet
	packetDown   gopacket.Packet
	anyPacket    gopacket.Packet
)

func Live_capture(dev, mac, protocol string) {
	// Open device
	handle, err = pcap.OpenLive(dev, snapshot_len, promiscuous, timeout)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	log.Println(protocol)
	// Set Filter
	if protocol != "a" {
		err = handle.SetBPFFilter(protocol)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Use the handle as a packet source to process all packets
	go viewData()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	/*for packet := range packetSource.Packets() {
		srcMac := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet).SrcMAC.String()
		if mac == srcMac {
			upBytes++ //len(packet.Data())
			//packetUp = packet
		} else {
			downBytes++ //len(packet.Data())
			//packetDown = packet
			nextPacket, _ := packetSource.NextPacket()
			//fmt.Println(packet.Metadata().Timestamp, nextPacket.Metadata().Timestamp)
			fmt.Print("\r", nextPacket.Metadata().Timestamp.Sub(packet.Metadata().Timestamp))
		}
		anyPacket = packet

	}*/

	chanPackets := packetSource.Packets()
	for {
		anyPacket = <-chanPackets
	}
}
