package main

import (
	"log"
	"github.com/google/gopacket/pcap"
	"fmt"
)

func main(){
	devices, err := pcap.FindAllDevs()
	if err != nil{
		log.Fatalf("Error reading devices : %s\n", err.Error())
	}
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\n Name: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Device Addresses: ")
		for _, address := range device.Addresses{
			fmt.Println("-IP Address : ", address.IP)
			fmt.Println("- Subtnet Address : ", address.Netmask)
		}
	}
}

