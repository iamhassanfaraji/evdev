package evdev

import (
  "os"
  "fmt"
)

// InputDevice is type devices that are connected to machine for user communication
type Device struct {
  Path string // path to input device (devnode)

  Name string   // device name
  Phys string   // physical topology of device
  File *os.File // an open file handle to the input device

  BusType uint16 // bus type identifier
  Vendor  uint16 // vendor identifier
  Product uint16 // product identifier
  Version uint16 // version identifier

  EvdevVersion int // evdev protocol version

  Capabilities []Capability
}

func (dev *Device) stringifyDevice() string {	

	return fmt.Sprintf(
		"InputDevice %s (fd %d)\n"+
			"  name %s\n"+
			"  phys %s\n"+
			"  bus 0x%04x, vendor 0x%04x, product 0x%04x, version 0x%04x\n"+
      dev.Path, dev.File.Fd(), dev.Name, dev.Phys, dev.BusType,
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
