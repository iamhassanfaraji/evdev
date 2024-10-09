package evdev

import (
	"strings"
)

type CapabilityType struct {
  name string
  code uint16
  max uint16
}

type CapabilityCode struct { 
  name string
  code uint16
}

type Capability struct{
  CapabilityType CapabilityType
  CapabilityCodes []CapabilityCode
}  

func findCodes(prefix string) []CapabilityCode{
  var codes []CapabilityCode
  for k,v := range inputCodes {
    if strings.HasPrefix(k, prefix){
      codes = append(codes, CapabilityCode{k, v})
    }
  } 
  
  return codes
}

func capabilityBuilder(capabilities *[]Capability, inputType string, prefixes ...string){
  var codes []CapabilityCode
  
  for _,v := range prefixes {
    codes = append(codes, findCodes(v)...)
  }

  *capabilities = append(*capabilities, 
    Capability{
        CapabilityType{
            inputType, 
            inputTypes[inputType],
            maxOfKeys[inputType],
        },
        codes, 
    },
  ) 
}

func generatePossibleCapabilities() []Capability {
  var possibleCapabilities []Capability
  
  for k,_ := range inputTypes{
    switch k {
        case "EV_SYN":
          capabilityBuilder(&possibleCapabilities, "EV_SYN", "SYN")
        case "EV_KEY":
          capabilityBuilder(&possibleCapabilities, "EV_KEY", "KEY","BTN")
        case "EV_REL":
          capabilityBuilder(&possibleCapabilities, "EV_REL", "REL")
        case "EV_ABS":
          capabilityBuilder(&possibleCapabilities, "EV_ABS", "ABS")
        case "EV_MSC":
          capabilityBuilder(&possibleCapabilities, "EV_MSC", "MSC")
        case "EV_SW":
          capabilityBuilder(&possibleCapabilities, "EV_SW", "SW")
        case "EV_LED": 
          capabilityBuilder(&possibleCapabilities, "EV_LED", "LED")
        case "EV_SND":
          capabilityBuilder(&possibleCapabilities, "EV_SND", "SND")
        case "EV_REP":
          capabilityBuilder(&possibleCapabilities, "EV_REP", "REP")
        case "EV_FF":
          capabilityBuilder(&possibleCapabilities, "EV_FF", "FF")
        case "EV_PWR":
          capabilityBuilder(&possibleCapabilities, "EV_PWR", "PWR")
    }
  } 

  return possibleCapabilities
}
