package _171bitand2bitCharacters

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	if isOneBitCharacter([]int{1,1,1,0}) {
		t.Error("can't pass")
	} else {
		t.Log("pass")
	}
}
