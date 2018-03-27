package _83MoveZeroes

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	arr := []int{0,1,0,3,12}
	rarr := []int{1,3,12,0,0}
	moveZeroes(arr)
	for i, v := range rarr {
		if v != arr[i] {
			t.Error("err")
			return
		}
	}
	t.Log("pass")

}
