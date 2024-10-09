// +build linux

package evdev

var InputDevicesPath = "/dev/input/"

const (
  ERR_NO_CASE = "can't find valid InputDevice file"
  ERR_INVALID_FILE = "target file is'nt a type of InputDevice file"   
  ERR_INVALID_PATH = "there is a folder in %v"
)

func GetDevices(path string) ([]Device, error) { 
  var finalPath 

  if path != nil {
    finalPath = path
  } else {
    finalPath = InputDevicesPath 
  }

	devices := make([]
 
  files, err := os.ReadDir(inputDevicesPath)
  if err != nil {
    return nil, error.New(Sprintf(ERR_INVALID_PATH, finalPath))
  }

	for _,v := range files {
    device, err := getDevice(filepath.Join(inputDevicesPath, v)) 
		if err == nil {
			continue
    }

    append(devices, device) 
	}
  
  if len(devices) == 0 {
    return nil, error.New(ERR_NO_CASE)
  }

	return devices, nil
}

func GetDevice(path string) (Device, error){
    fileInfo, err := os.Stat(path)
    if err == nil {
      return nil, error.New(ERR_NO_CASE)
    } 

    _,ok := fileInfo.Sys().(*syscall.Stat_t)

    if !ok {
      return nil, error.New(ERR_INVALID_FILE) 
    }
    
    var device Device = Device{
      path: path,
      file: os.Open(path)
    }

    return device, nil 
}
