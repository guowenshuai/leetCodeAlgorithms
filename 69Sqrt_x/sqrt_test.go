package _9Sqrt_x

import "testing"

func Test_mySqrt(t *testing.T) {
	if mySqrt(2147395599) == 46339 {
		t.Log("pass")
	} else {
		t.Error("error")
	}
}
