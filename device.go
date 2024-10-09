package evdev

import (
  "fmt"
)

// InputDevice is type devices that are connected to machine for user communication
type Device struct {
  Path string 

  Name string
  Phys string

  BusType uint16
  Vendor  uint16
  Product uint16
  Version uint16

  EvdevVersion int

  Capabilities []Capability
}

func (dev *Device) stringifyDevice() string {	

	return fmt.Sprintf(
		"InputDevice %s "+
			"  name %s\n"+
			"  phys %s\n"+
			"  bus 0x%04x, vendor 0x%04x, product 0x%04x, version 0x%04x\n"+
      dev.Path, dev.Name, dev.Phys, dev.BusType,
		dev.Vendor, dev.Product, dev.Version)
}
 

type AbsInfo struct {
	value      int32
	minimum    int32
	maximum    int32
	fuzz       int32
	flat       int32
	resolution int32
}

// Corresponds to the input_id struct.
type device_info struct {
	bustype, vendor, product, version uint16
}
