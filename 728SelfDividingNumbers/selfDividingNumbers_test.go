package _28SelfDividingNumbers

import "testing"

func Test_selfDividingNumbers(t *testing.T) {
	tmp := selfDividingNumbers(1, 22)
	if len(tmp) != 13{
		t.Error("test error")
	} else {
		t.Log("pass")
	}
}
