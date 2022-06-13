package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ecc1/ble"
)

var uuid string = "6e400001-b5a3-f393-e0a9-e50e24dcca9e"

func main() {
	count := 0
	for {
		count++
		log.Printf("count=%d\r\n", count)
		conn, err := ble.Open()
		if err != nil {
			log.Fatal(err)
		}
		//conn.Print(os.Stdout)

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
		time.Sleep(time.Second * 3)

		device.Print(os.Stdout)

		//去断开
		err = device.Disconnect()
		if err != nil {
			log.Fatal(err)
			//log.Println(err)
		}
		device.Print(os.Stdout)

		if err != nil {
			log.Fatal(err)
		}

		//去冲洗
		adapter, err := conn.GetAdapter()
		adapter.Print(os.Stdout)
		adapter.RemoveDevice(device)

		time.Sleep(time.Second)
	}
}
func ResetBle() {

	cmd := exec.Command("/etc/init.d/bluetooth", "restart")
	stdout, err := cmd.Output()
	if err != nil {
		log.Printf("[ResetBle]exec.Command fail %v\r\n", err)
	} else {
		log.Printf("[ResetBle]exec.Command ok %s\r\n", stdout)
	}

}
func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
