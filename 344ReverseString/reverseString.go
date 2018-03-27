package _44ReverseString

import (
)

func reverseString(s string) string {
	sBytes := []byte(s)
	for i,j:=0,len(sBytes)-1; i<j; i++ {
		sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
		j--
	}
	return string(sBytes)
}