package _76NumberComplement

import "testing"

func Test_findComplement(t *testing.T){
	//t.Log(findComplement(5))
	if findComplement(1111) != 936 {
		t.Error("can't pass")
	} else {
		t.Log("pass")
	}
}
