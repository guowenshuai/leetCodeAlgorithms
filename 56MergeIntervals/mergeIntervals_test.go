package _6MergeIntervals

import "testing"

func Test_merge(t *testing.T) {

	t.Log(merge([]Interval{Interval{0,2}, Interval{2,3}, Interval{4,4}, Interval{0,1},
	Interval{5,7}, Interval{4,5}, Interval{0,0}}))



}
