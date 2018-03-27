package _75DistributeCandies

import "testing"

func Test_distributeCandies(t *testing.T) {
	if distributeCandies([]int{1,1,2,2,3,3,}) == 3 {
		t.Log("pass")
	} else {
		t.Error("can]t pass")
	}
}