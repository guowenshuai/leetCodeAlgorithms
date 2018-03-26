package _87RepeatedDNASequences

import (
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	s := "AAAAAAAAAAA"
	if r := findRepeatedDnaSequences(s); len(r) != 1 {
		t.Error("can't psss")
	} else {
		t.Log("pass")
	}
}