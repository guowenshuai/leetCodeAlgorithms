package _87RepeatedDNASequences

import (
	"strings"
)


func findRepeatedDnaSequences(s string) []string {
	regMap := make(map[string]int)
	ret := []string{}
	if len(s) < 11 {
		return ret
	}
	sByte := []byte(s)
	sLen := len(s)
	for i:=0; i<sLen-9; i++ {
		index := string(sByte[i:i+10])
		if _, ok := regMap[index]; ok {
			if(regMap[index] == 1) {
				regMap[index]++
				ret = append(ret, index)
			}
		} else {
			regMap[index] = 1
		}
	}
	return ret
}

func findRepeatedDnaSequencesBK(s string) []string {
	sTmp := []byte(s)
	regMap := make(map[string]int)
	ret := []string{}
	for i:=0; i<len(s)-10;i++ {
		if _, ok := regMap[string(sTmp[i : i+10])]; !ok {
			if strings.Contains(string(sTmp[i+1:len(s)]), string(sTmp[i:i+10])) {
				regMap[string(sTmp[i : i+10])] = 0
				ret = append(ret, string(sTmp[i : i+10]))
				}
				}
				}
	return ret
}