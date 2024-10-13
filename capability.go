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
  capabilityType CapabilityType
  capabilityCodes []CapabilityCode
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


func capabilitySeter(id IDDevice) []Capability{
  var possibleCapabilities []Capability = generatePossibleCapabilities()
 
  var availableCapabilities []Capability
  
  for k,v := range possibleCapabilities{
    ok, err := v.capabilityType.checkInputType(id)
    if ok {
      availableCapabilities = append(availableCapabilities, Capability{})
      availableCapabilities[k].capabilityType = v.capabilityType  
      for _, vc := range v.capabilityCodes {
        ok, err := vc.checkInputCode(id, v.capabilityType)
        if ok {
           availableCapabilities[k].capabilityCodes = append(availableCapabilities[k].capabilityCodes, vc) 
        } else if !ok && err == nil {
          continue
        }else{
          panic(err)
        }
      }
    } else if !ok && err == nil{
      continue
    } else {
      panic(err)
    } 
  } 

  return availableCapabilities
}

func CapabilityTypeProvider (f func(string, uint16, uint16) bool, arg *CapabilityType) bool {
  return f(arg.name, arg.code, arg.max) 
}

func CapabilityCodeProvider (f func(string, uint16) bool, arg *CapabilityCode) bool {
  return f(arg.name, arg.code)
}

