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

  name Name
  phys Phys

  busType uint16
  vendor  uint16
  product uint16
  version uint16

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

func GetDevice (id IDDevice) (Device, error){
  name, phys, deviceInfo, err := deviceInfoExtractor(id) 
  if err != nil {
    return Device{}, err 
  }

  return Device{
    id,
    name,
    phys,
    deviceInfo.busType,
    deviceInfo.vendor,
    deviceInfo.product,
    deviceInfo.version,    
    capabilitySeter(id),
  }, nil
}

func GetDevices() (*[]Device, error){
  ids, err := availableDevices() 

  if err != nil{
    return nil, err   
  }

  var devices []Device 

  for _,id := range *ids {
    device, err := GetDevice(id)
    if err != nil {
      return nil, err
    }
    devices = append(devices, device)
  }

  return &devices, nil
}
