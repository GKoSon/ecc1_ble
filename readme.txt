准备ESP32C3 固件
3根线
修改TX+18db(不修改找不到从机的)
sudo btattach -N -B /dev/ttyUSB0 -S 921600

1---选择HCI1 怎么操作?
在获得的时候 增加判断条件
func (conn *Connection) GetAdapter() (Adapter, error) {
	return conn.findObject(adapterInterface, func(koson *blob) bool {
		if koson.path == "/org/bluez/hci1" {
			return true
		}
		return false
	})
}

2---增加名字的过滤 怎么操作?
const target_name = "M_IZAR_SH1"

func (adapter *blob) SetDiscoveryFilter(uuids ...string) error {
	log.Printf("%s: setting discovery filter %v", adapter.Name(), UUIDs(uuids))
	return adapter.call(
		"SetDiscoveryFilter",
		Properties{
			"Transport": dbus.MakeVariant("le"),
			"UUIDs":     dbus.MakeVariant(uuids),
			"Pattern":   dbus.MakeVariant(target_name),
		},
	)
}
3---整理自己的mcube流程
第二次就会失败
找到2个设备 没有移除掉吗?