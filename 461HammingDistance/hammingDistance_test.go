package _61HammingDistance

import "testing"

func Test_hammingDistance(t *testing.T) {
	if num := hammingDistance(1, 4); num != 2 {
		t.Error("error hammingDistance(1, 4): ", num)
	} else {
		t.Log("test success")
	}
}
