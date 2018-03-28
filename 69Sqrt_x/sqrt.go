package _9Sqrt_x


func mySqrt(x int) int {
	zeros := -1
	r := 0
	for tmp:=x; tmp>0; tmp >>= 1{
		zeros++
	}
	if zeros%2 == 0 {
		r = 1<<uint(zeros/2)
	} else {
		r = 1<<uint((zeros+1)/2)
	}

	if r*r >= x {
		for r*r > x {
			r--
		}
		return r
	} else if r*r < x {
		for r*r < x {
			r++
		}
		return r-1
	}

	return r
}

/*
1	01		01
2	10		100
3	11		1001
4	100		10000
5	101		11001
6	110		100100
7	111		110001	->	1000000	->	1000	->	7
8	1000	1000000
*/