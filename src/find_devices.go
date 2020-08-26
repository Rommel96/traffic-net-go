package src

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/google/gopacket/pcap"
)

type NIC struct {
	Name   string
	Device string
	Ips    []string
	Mac    string
}

func Find_devices() []NIC {
	// Find all devices
	var nics []NIC
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	itf, _ := net.Interfaces()

	// Print device information
	fmt.Println("Devices found:")
	for index, i := range itf {
		nic := NIC{
			Name: i.Name,
			Mac:  i.HardwareAddr.String(),
		}
		nics = append(nics, nic)
		ips, _ := i.Addrs()
		for _, ip := range ips {
			for _, device := range devices {
				for _, addrs := range device.Addresses {
					if strings.Contains(ip.String(), addrs.IP.String()) {
						nics[index].Ips = append(nics[index].Ips, ip.String())
						nics[index].Device = device.Name
					}
				}
			}
		}
	}

	return nics

}
