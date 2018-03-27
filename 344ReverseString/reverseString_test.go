package _44ReverseString

import "testing"

func Test_reverseString(t *testing.T) {
	if reverseString("hello") != "olleh" {
		t.Error("error")
	} else {
		t.Log("pass")
	}
}
