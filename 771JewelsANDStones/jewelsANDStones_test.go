package _71JewelsANDStones

import (
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	if r, e := numJewelsInStones("aA", "aAAkdwSwed"); r != 3 || e != nil {
		//t.Log(r)
		t.Error("error test")
	} else {
		t.Log("test pass")
	}
}

