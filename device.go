package evdev

import (
  "fmt"
)

type IDDevice string
type Name string
type Phys string

// InputDevice is type devices that are connected to machine for user communication
// field of a Device should be private, any changes able with method
type Device struct {
  id  IDDevice// in linux initialized by path, in macos initialized by hardware UUID , in windows initialized by hardware ID

  name string
  phys string

  busType uint16
  vendor  uint16
  product uint16
  version uint16

  evdevVersion int

  capabilities []Capability
}

func (dev *Device) StringifyDevice() string {	

	return fmt.Sprintf(
		"InputDevice %s "+
			"  name %s\n"+
			"  phys %s\n"+
			"  bus 0x%04x, vendor 0x%04x, product 0x%04x, version 0x%04x\n"+
      string(dev.id), dev.name, dev.phys, dev.busType,
		dev.vendor, dev.product, dev.version)
}

type DeviceInfo struct {
    busType uint16
    vendor  uint16
    product uint16
    version uint16
}
