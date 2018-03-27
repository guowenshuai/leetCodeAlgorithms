package _5SearchInsertPosition

import "testing"

func Test_searchInsert(t *testing.T){
	if searchInsert([]int{1,2,5,6}, 5) != 2 {
		t.Error("can't pass")
	} else {
		t.Log("pass")
	}
}



