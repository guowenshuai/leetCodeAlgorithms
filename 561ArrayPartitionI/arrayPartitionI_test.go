package _61ArrayPartitionI

import "testing"

func Test_arrayPairSum(t *testing.T) {
	if arrayPairSum([]int{1,3,4,2}) == 4 {
		t.Log("pass")
	} else {
		t.Error("can't pass")
	}
}

func Test_qSort(t *testing.T) {
	arr := []int{3381,
		443,
		6918,
		-4429,
		1531,
		4315,
		862,
		-206,2,2}
	qSort(arr)
	t.Log(arr)


}