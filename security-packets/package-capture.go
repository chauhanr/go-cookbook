package main

import (
	"time"
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket"
	"fmt"
)

var (
	device = "eth0"
	snapshotLen int32 = 1024
	promiscuous  = false
	err error
	timeout = 30 * time.Second
	handle *pcap.Handle
)

func main(){
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil{
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// process package
		fmt.Println(packet)
	}

}
