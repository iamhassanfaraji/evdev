package evdev

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
  "error"
)

// InputDevice is type devices that are connected to machine for user communication
type Device struct {
  path string // path to input device (devnode)

  name string   // device name
  phys string   // physical topology of device
  file *os.File // an open file handle to the input device

  bustype uint16 // bus type identifier
  vendor  uint16 // vendor identifier
  product uint16 // product identifier
  version uint16 // version identifier

  evdevVersion int // evdev protocol version

  capabilities     map[CapabilityType][]CapabilityCode // supported event types and codes.
  capabilitiesFlat map[int][]int
}

// Read and return a slice of input events from device.
func (input *Device) getInputEvents() ([]InputEvent, error) {
	events := make([]InputEvent, 16)
	buffer := make([]byte, eventsize*16)

	_, err := dev.File.Read(buffer)
	if err != nil {
		return events, err
	}

	b := bytes.NewBuffer(buffer)
	err = binary.Read(b, binary.LittleEndian, &events)
	if err != nil {
		return events, err
	}

	// remove trailing structures
	for i := range events {
		if events[i].Time.Sec == 0 {
			events = append(events[:i])
			break
		}
	}

	return events, err
}

// Read and return a single input event.
func (dev *Device) getInputEvent() (*InputEvent, error) {
	event := InputEvent{}
	buffer := make([]byte, eventsize)

	_, err := dev.File.Read(buffer)
	if err != nil {
		return &event, err
	}

	b := bytes.NewBuffer(buffer)
	err = binary.Read(b, binary.LittleEndian, &event)
	if err != nil {
		return &event, err
	}

	return &event, err
}

func (dev *Device) stringifyDevice() string {
	evtypes := make([]string, 0)

	for ev := range dev.Capabilities {
		evtypes = append(evtypes, fmt.Sprintf("%s %d", ev.Name, ev.Type))
	}
	evtypes_s := strings.Join(evtypes, ", ")

	return fmt.Sprintf(
		"InputDevice %s (fd %d)\n"+
			"  name %s\n"+
			"  phys %s\n"+
			"  bus 0x%04x, vendor 0x%04x, product 0x%04x, version 0x%04x\n"+
			"  events %s",
		dev.Fn, dev.File.Fd(), dev.Name, dev.Phys, dev.Bustype,
		dev.Vendor, dev.Product, dev.Version, evtypes_s)
}


// An all-in-one function for describing an input device.
func (dev *Device) set_device_info() error {
	info := device_info{}

	name := new([MAX_NAME_SIZE]byte)
	phys := new([MAX_NAME_SIZE]byte)

	err := ioctl(dev.File.Fd(), uintptr(EVIOCGID), unsafe.Pointer(&info))
	if err != 0 {
		return err
	}

	err = ioctl(dev.File.Fd(), uintptr(EVIOCGNAME), unsafe.Pointer(name))
	if err != 0 {
		return err
	}

	// it's ok if the topology info is not available
	ioctl(dev.File.Fd(), uintptr(EVIOCGPHYS), unsafe.Pointer(phys))

	dev.Name = bytes_to_string(name)
	dev.Phys = bytes_to_string(phys)

	dev.Vendor = info.vendor
	dev.Bustype = info.bustype
	dev.Product = info.product
	dev.Version = info.version

	ev_version := new(int)
	err = ioctl(dev.File.Fd(), uintptr(EVIOCGVERSION), unsafe.Pointer(ev_version))
	if err != 0 {
		return err
	}
	dev.EvdevVersion = *ev_version

	return nil
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
