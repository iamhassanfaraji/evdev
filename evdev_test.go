package evdev

import (
  "testing" 
)

func TestGeneratorCapabilities(t *testing.T){
  var expectation = map[string]int{
    "EV_SYN": 4,
	  "EV_KEY": 559,
	  "EV_REL": 11,
	  "EV_ABS": 42,
	  "EV_MSC": 6,
	  "EV_SW":  17,
	  "EV_LED": 11,
	  "EV_SND": 3,
	  "EV_REP": 2,
	  "EV_FF":  24,
    "EV_PWR": 0,
  }

  capabilities := generatePossibleCapabilities() 
  var unexpectedCapabilityCodesNumber = make(map[string]int) 

  for _,v := range capabilities {
    if expectation[v.capabilityType.name] == len(v.capabilityCodes){
      continue
    }else{
      unexpectedCapabilityCodesNumber[v.capabilityType.name] = len(v.capabilityCodes)
    }
  } 
 
  if len(unexpectedCapabilityCodesNumber) != 0 { 
    t.Error("there is'nt build capabilities in right range") 
  } 
}
