package _06NumberofLinesToWriteString

import (
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	widths := []int{10,4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	S := "aaaaaaaaabaaaaaaaaabaaaaaa"
	if ret := numberOfLines(widths, S); ret[0]!= 3 || ret[1] != 60 {
		t.Error("error")
	} else {
		t.Log("pass")
	}
}
