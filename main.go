package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rommel96/traffic-net-go/src"
)

var (
	colorRed   = "\033[31m"
	colorReset = "\033[0m"
)

func main() {
	nics := src.Find_devices()

	for i, nic := range nics {
		fmt.Printf("[%d]: ", i)
		fmt.Println("	Name: ", nic.Name)
		fmt.Println("	Device: ", nic.Device)
		fmt.Println("	IPs: ", nic.Ips)
		fmt.Println("	Mac: ", nic.Mac)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Choose NIC[0-%d, default 0]: ", len(nics)-1)
		textInput, _ := reader.ReadString('\n')
		textInput = strings.TrimRight(textInput, "\r\n")
		indexInterface, _ := strconv.Atoi(textInput)
		if indexInterface < 0 || indexInterface > len(nics) {
			fmt.Println(string(colorRed), "Option incorrect. Try again.", string(colorReset))
		} else {
			reader = bufio.NewReader(os.Stdin)
			fmt.Print("Entry Filter[a for All]: ")
			protocol, _ := reader.ReadString('\n')
			protocol = strings.TrimRight(protocol, "\r\n")
			src.Live_capture(nics[indexInterface].Device, nics[indexInterface].Mac, strings.ToLower(protocol))
			break
		}
	}
}
