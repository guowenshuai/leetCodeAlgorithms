package _28SelfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	if left < 1 {
		left = 1
	}
	if right > 10000 {
		right = 10000
	}

	for ;left<=right;left++ {
		if isSelfDivided(left) {
			ret = append(ret, left)
		}
	}
	return ret
}

func isSelfDivided(num int) bool {
	tmp := num
	d := 1
	for tmp>0 {
		if(tmp >= 10) {
			d = tmp%10
			tmp /= 10
		} else {
			d = tmp
			tmp = -1
		}
		if d == 0 {
			return false
		}
		if num%d != 0 {
			return false
		}
	}
	return true
}