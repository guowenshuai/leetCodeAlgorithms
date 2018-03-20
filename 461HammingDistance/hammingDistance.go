package _61HammingDistance

import (
	"math"
	"fmt"
)

func hammingDistance(x int, y int) int {
	if x < 0 || y > int(math.Pow(2, 31)){
		return 0
	}
	r := x ^ y
	fmt.Println("r:", r)
	count := 0
	for r > 0 {
		count++
		r &= r-1
	}
	return count
}
