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
