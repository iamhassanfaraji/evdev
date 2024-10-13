// +build linux

package evdev 

import (
  "fmt"
  "errors"
  "syscall"
  "unsafe"
  "os"
  "path/filepath"
)

var InputDevicesPath = "/dev/input/"

const (
  ERR_NO_CASE = "can't find valid InputDevice file"
  ERR_INVALID_FILE = "target file is'nt a type of InputDevice file"   
  ERR_INVALID_PATH = "there is a folder in %v"
  ERR_OPEN_FILE = "in loading input device(id: %v), get this error: %v"
  ERR_NAME = "can't resolve name of device (id: %v)"
  ERR_PHYS = "can't resolve phys of device(id: %v)"
  ERR_INPUT_ID = "can't resolve inputId of device(id: %v)"
  ERR_CHECK_INPUT_TYPE = "can't resolve availablity of input type(%v) of device(id: %v)"
  ERR_CHECK_INPUT_CODE = "can't resolve availablity of input code(%v) of input type(%v) of device(%v)"
)

func availableDevices() (*[]IDDevice, error) {  
	var ids []IDDevice
 
  files, err := os.ReadDir(InputDevicesPath)
  if err != nil {
    return nil, errors.New(fmt.Sprintf(ERR_INVALID_PATH, InputDevicesPath))
  }

	for _,v := range files {
    ids = append(ids, IDDevice(filepath.Join(InputDevicesPath, v.Name()))) 
	} 

	return &ids, nil
}

func deviceInfoExtractor(path IDDevice) (Name, Phys, DeviceInfo, error){
    file, err := os.Open(string(path))
    if err == nil {
      return "", "", DeviceInfo{}, err
    } 

    // check the id(file) is type of input device
    fileInfo, infoError := file.Stat()
    if infoError != nil {
      return "", "", DeviceInfo{}, errors.New(ERR_INVALID_FILE)
    }

    _,ok := fileInfo.Sys().(*syscall.Stat_t)
    if !ok {
      return "", "", DeviceInfo{}, errors.New(ERR_INVALID_FILE) 
    }

    name := make([]byte, 256)
    nerr := ioctl(file.Fd(), EVIOCGNAME, unsafe.Pointer(&name[0]))
    if nerr != 0 {
        return "", "", DeviceInfo{}, errors.New(fmt.Sprintf(ERR_NAME, path))
    }


    phys := make([]byte, 256)
    perr := ioctl(file.Fd(), EVIOCGPHYS, unsafe.Pointer(&phys[0]))
    if perr != 0 {
        return "", "", DeviceInfo{}, errors.New(fmt.Sprintf(ERR_PHYS, path))     
    } 

    var info DeviceInfo
    ierr := ioctl(file.Fd(), EVIOCGID, unsafe.Pointer(&info))
    if ierr != 0 {
        return "", "", DeviceInfo{}, errors.New(fmt.Sprintf(ERR_INPUT_ID, path)) 
    }

    file.Close() 
    return Name(name), Phys(phys), DeviceInfo(info), nil 
}  

func ioctl(fd uintptr, request uintptr, arg unsafe.Pointer) int {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, request, uintptr(arg))
	if errno != 0 {
		return int(errno) 
	}
	return 0
}

func (inputType CapabilityType) checkInputType (path IDDevice) (bool, error){
  file,_ := os.Open(string(path))
  fd := file.Fd()
  defer file.Close()
  
  evBits := make([]byte, (inputType.max+7)/8)

	err := ioctl(fd, EVIOCGBIT, unsafe.Pointer(&evBits[0]))
	if err != 0 {
    return false, errors.New(fmt.Sprintf(ERR_CHECK_INPUT_TYPE, inputType.name, string(path)))
	}

	if evBits[inputType.code/8]&(1<<(inputType.code%8)) != 0 {
		return true, nil
	}

	return false, nil 
}

func (inputCode CapabilityCode) checkInputCode (path IDDevice, inputType CapabilityType) (bool, error){
  file,_ := os.Open(string(path))
  fd := file.Fd()
  defer file.Close() 

	codeBits := make([]byte, (inputType.max+7)/8)

	err := ioctl(fd, EVIOCGBIT|uintptr(inputType.code)<<8, unsafe.Pointer(&codeBits[0]))
	if err != 0 {
	  return false, errors.New(fmt.Sprintf(ERR_CHECK_INPUT_CODE, inputCode.name, inputType.name, string(path)))	
	}

	if codeBits[inputCode.code/8]&(1<<(inputCode.code%8)) != 0 {
		return true, nil
	}

	return false, nil
}
