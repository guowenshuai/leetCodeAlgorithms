package _76NumberComplement

import "fmt"

func findComplement(num int) int {
	fmt.Printf("%b\n", num)

	r := num ^ 0xffffffff
	fmt.Printf("%b\n", r)
	tool := 0
	for num>0 {
		num >>= 1
		if tool > 0 {
			tool <<= 1
		}
		tool++

	}
	fmt.Printf("%b\n", tool)

	return r&tool
}
//
//
// 00000101
//^
// 11111010

//&00000111
// 00000010
