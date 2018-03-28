package _6MergeIntervals

import (
	"sort"
	"fmt"
)

//Definition for an interval.
type Interval struct {
	Start int
	End   int
}
type data []Interval
func (d data) Len() int {
	return len(d)
}
func (d data) Less (i, j int) bool {
	return d[i].Start < d[j].Start
}
func (d data) Swap (i, j int) {
	d[i], d[j] = d[j], d[i]
}
func merge(intervals []Interval) []Interval {
	sort.Sort(data(intervals))
	fmt.Println(intervals)
	if len(intervals) == 0 {
		return []Interval{}
	}
	ret := []Interval{intervals[0]}
	for _, v := range intervals {
		insert := false
		for j, r := range ret {
			if v.Start >= r.Start && v.Start <= r.End {
				insert = true
				if v.End > r.End {
					ret[j].End = v.End
					break
				}
			} else if v.Start < r.Start && v.End >= r.Start {
				insert = true
				ret[j].Start = v.Start
				if v.End > r.End {
					ret[j].End = v.End
					break
				}
			}
		}
		if insert {
			continue
		} else {
			ret = append(ret, v)
		}

	}
	return ret
}
