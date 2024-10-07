package evdev

import (
	"fmt"
	"syscall"
	"unsafe"
)

type InputEvent struct{
  Time  syscall.Timeval 
  Type  uint16          	
  Code  uint16          
  Value int32 
}


// human readible a event
func (ev *InputEvent) stringify() string {
	return fmt.Sprintf("event at %d.%d, code %02d, type %02d, val %02d",
		ev.Time.Sec, ev.Time.Usec, ev.Code, ev.Type, ev.Value)
}

