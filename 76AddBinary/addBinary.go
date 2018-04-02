package _6AddBinary

import (
	"fmt"
)

func addBinary(a string, b string) string {
	tmpa, tmpb := []byte(a), []byte(b)
	ret := []byte{}
	t := 0
	for i,j := len(a)-1,len(b)-1; i>-1 || j>-1; {
		x,y := 0,0
		if i >= 0 {
			x = int(tmpa[i])-48
		}
		if j>= 0 {
			y = int(tmpb[j])-48
		}
		//u,v := add(x, y ,t)
		switch x+y+t {
		case 0:
			t = 0
			ret = append(ret, byte(0+48))
		case 1:
			t = 0
			ret = append(ret, byte(1+48))
		case 2:
			t = 1
			ret = append(ret, byte(0+48))
		case 3:
			t = 1
			ret = append(ret, byte(1+48))
		}
		i--
		j--
	}
	if t == 1 {
		ret = append(ret, byte(1+48))
	}
	for i,j:=0,len(ret)-1; i<j; i++ {
		ret[i], ret[j] = ret[j], ret[i]
		j--
	}
	fmt.Println(string(ret))

	return string(ret)

}