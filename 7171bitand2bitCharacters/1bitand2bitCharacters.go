package _171bitand2bitCharacters

func isOneBitCharacter(bits []int) bool {
	i:=0
	for i<len(bits)-1 {
		i += bits[i]+1
	}
	return i == len(bits)-1
}
