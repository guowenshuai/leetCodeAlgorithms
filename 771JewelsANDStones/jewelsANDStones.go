package _71JewelsANDStones

import (
	"errors"
	"strings"
)

func numJewelsInStones(J string, S string) (count int, e error) {

	if len(J) > 50 || len(S) > 50{
		return count, errors.New("len of string more than 50")
	}


	for _, c := range S {
		if strings.Contains(J, string(c)) {
			count++
		}
	}

	return count, e
}
