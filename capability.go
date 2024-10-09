package evdev

import (
	"strings"
)

type CapabilityType struct {
  Name string
  Code uint16
}

type CapabilityCode struct { 
  Name string
  Code uint16
}

type Capability struct{
  CapabilityType CapabilityType
  CapabilityCodes []CapabilityCode
} 

func checkPrefix(value string, prefixes string) bool{
  var speratedPrefix = strings.Split(prefixes, ",") 
  for i := 0; i < len(speratedPrefix); i++ {

    if strings.HasPrefix(value, speratedPrefix[i]) {
      return true
    } else {
      continue
    }
  }  

  return false
}

func generatePossibleCapabilities() []Capability {
  var possibleCapabilities []Capability

  for inputTypeName, inputTypeValue := range inputTypes {
    var capabilityCodes []CapabilityCode
    
    for inputCodeName, inputCodeValue := range inputCodes{
      
      if checkPrefix(inputCodeName, prefixCodeOfInputTypes[inputTypeValue]) {
        capabilityCodes = append(
          capabilityCodes, 
          CapabilityCode{
            inputCodeName,
            inputCodeValue,
          },
        )
      }else{
        continue
      }
    }  

    possibleCapabilities = append(
      possibleCapabilities, 
      Capability{
        CapabilityType{inputTypeName, inputTypeValue},
        capabilityCodes,
      },
    )
  }

  return possibleCapabilities
}
