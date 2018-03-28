package _57ReverseWordsInAStirngIII

import "strings"

func reverseWords(s string) string {
	var reverseWord func(string) string
	reverseWord = func(s string) string{
		sBytes := []byte(s)
		for i,j:=0,len(sBytes)-1; i<j; i++ {
			sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
			j--
		}
		return string(sBytes)
	}

	words := strings.Split(s, " ")
	ret := []string{}
	for _, w := range words {
		ret = append(ret, reverseWord(w))
	}

	return strings.Join(ret, " ")
}