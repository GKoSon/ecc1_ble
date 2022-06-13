// +build !nofilter

package ble

import (
	"log"

	"github.com/godbus/dbus"
)

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
