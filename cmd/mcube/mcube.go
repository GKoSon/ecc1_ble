package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ecc1/ble"
)

var uuid string = "6e400001-b5a3-f393-e0a9-e50e24dcca9e"

func main2() {

	conn, err := ble.Open()
	if err != nil {
		log.Fatal(err)
	}
	conn.Print(os.Stdout)

	conn.RmObject("org.bluez.Device1") //一句话 全部冲洗
	time.Sleep(time.Second * 1)
	err = conn.Update()
	if err != nil {
		log.Fatal(err)
	}
	ShowTree()
}

func main() {
	count := 0

	for {
		count++
		log.Printf(" ============ [count=%d] ============ \r\n", count)
		conn, err := ble.Open()
		if err != nil {
			log.Fatal(err)
		}
		//conn.Print(os.Stdout)
		//获得适配器 + 获得设备 方式1
		adapter, err := conn.GetAdapter()
		if err != nil {
			log.Fatal(err)
		}
	LOOP:
		log.Printf(" ============ Discover ============ \r\n")
		err = adapter.Discover(time.Second, uuid)
		if err != nil {
			goto LOOP
			log.Fatal(err)
		}

		log.Printf(" ============ Update ============ \r\n")
		err = conn.Update()
		if err != nil {
			log.Fatal(err)
		}

		device, err := conn.GetDeviceByUUID(uuid)

		//获得设备 方式1的打包即这句话
		//	device, err := conn.Discover(0, "", uuid)
		if err != nil {
			log.Fatal(err)
		}
		device.Print(os.Stdout)

		log.Printf(" ============ device %s ============ \r\n", device.ShortAddress())
		//去连接
		if !device.Connected() {
			err = device.Connect()
			if err != nil {
				//log.Fatal(err)
				adapter.RemoveDevice(device)
				log.Println(err)
				goto LOOP
			}
		} else {
			log.Printf("%s: already connected", device.Name())
		}
		time.Sleep(time.Second * 3)

		device.Print(os.Stdout)
		log.Printf(" ============ device %v ============ \r\n", device.Connected())
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
		//adapter.Print(os.Stdout)
		//adapter.RemoveDevice(device)
		conn.RmObject("org.bluez.Device1") //一句话 全部冲洗
		time.Sleep(time.Second * 5)
		ShowTree()

	}
}

func ShowTree() {

	cmd := exec.Command("busctl", "tree", "org.bluez")
	stdout, err := cmd.Output()
	if err != nil {
		log.Printf("[ResetBle]exec.Command fail %v\r\n", err)
	} else {
		log.Printf("[ResetBle]exec.Command ok %s\r\n", stdout)
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

//DF:5F:69:77:55:2F -> DF5F6977552F
func String_rm_char(a string, b string) string {
	mac := ""
	str := strings.Split(a, b)
	for _, s := range str {
		mac += s
	}
	return mac
}
