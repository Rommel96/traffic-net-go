package src

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rommel96/traffic-net-go/src/server"
)

func saveAsCSV() {
	f, err := os.OpenFile("packetsFromChannel.csv", os.O_CREATE, 0755) // For write access
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//f.WriteString(fmt.Sprintf("time,upBytes,downBytes\n"))
	for {
		//f.WriteString(fmt.Sprintf("%s,%s,%s\n", time.Now().Format(time.StampMilli), packetUp, packetDown))
		//os.Stdout.WriteString(fmt.Sprintf("\rDown:%.2fkb/s \t Up:%dMbps ", float32(downBytes)/1024/1, upBytes))
		if anyPacket != nil {
			f.WriteString(string(anyPacket.NetworkLayer().LayerType().LayerTypes()[0]))
		}
		upBytes = 0
		downBytes = 0
		//log.Println(packetUp.Layers())
		//time.Sleep(1 * time.Second)
	}
}

func sendWS() {
	go server.Start()
	for {
		server.Broadcast <- server.Message{
			DownBytes: downBytes,
			UpBytes:   upBytes,
		}
		upBytes = 0
		downBytes = 0
		time.Sleep(500 * time.Millisecond)
	}
}

func viewData() {
	for {
		//os.Stdout.WriteString(fmt.Sprintf("\rDown:%d  Up:%d", downBytes, upBytes))
		if anyPacket != nil {
			os.Stdout.WriteString(fmt.Sprintf("%s", anyPacket))
		}
		upBytes = 0
		downBytes = 0
		anyPacket = nil
		time.Sleep(1 * time.Second)
	}
}
