package evdev
import ("testing")

type CapabilityType struct {
  name string
  code uint16
}

type CapabilityCode struct { 
  name string
  code uint16
}

type Capability struct{
  capabilityType CapabilityType
  capabilityCodes []CapabilityCode
} 

var possibleCapabilities []Capability

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

func generatePossibleCapabilities() {
  for inputTypeName, inputTypeValue := range inputType {
    var capabilityCodes []CapabilityCode
    
    for inputCodeName, inputCodeValue := range inputCodes{
      
      if checkPrefix(inputCodeName, prefixCodeOfInputTypes[inputTypeValue]) {
        append(capabilityCodes, CapabilityCode{inputCodeName, inputCodeValue}  
      }else{
        continue
      }
    }  

    append(possibleCapabilities, Capability{
      CapabilityType{inputTypeName, inputTypeValue},
      capabilityCodes,
    } 
  }
}

func TestGeneratorCapabilities(t *test.T){
  t.Error("fake error")  
}

func init(){
  generatePossibleCapabilities()  
}
