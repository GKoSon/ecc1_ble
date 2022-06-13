package main

import (
	"log"
	"os"
	"time"

	"github.com/ecc1/ble"
)

var uuid string = "6e400001-b5a3-f393-e0a9-e50e24dcca9e"

func main() {
	conn, err := ble.Open()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for {
		count++
		log.Printf("count=%d\r\n", count)

		device, err := conn.Discover(0, "", uuid)
		if err != nil {
			log.Fatal(err)
		}
		device.Print(os.Stdout)

		//去连接

		if !device.Connected() {
			err = device.Connect()
			if err != nil {
				log.Fatal(err)
				//log.Println(err)
			}
		} else {
			log.Printf("%s: already connected", device.Name())
		}

		//去断开
		err = device.Disconnect()
		if err != nil {
			log.Fatal(err)
			//log.Println(err)
		}

		adapter, err := conn.GetAdapter()
		if err != nil {
			log.Fatal(err)
		}
		//去冲洗
		//adapter.Print(os.Stdout)
		adapter.RemoveDevice(device)
		time.Sleep(time.Second)
	}
}
