package evdev

import (
	"fmt"
	"strings"
  "testing"
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
  for _,p := range speratedPrefix {
    var result= strings.HasPrefix(value, p) 
    if result == true {
      continue
    } else {
      return false
    }
  }  

  return true
}

func inspect(data []Capability){
  var err = recover()
  fmt.Println(err)
  fmt.Println(data)
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

func TestGeneratorCapabilities(t *testing.T){
  t.Error("fake error")  
}
