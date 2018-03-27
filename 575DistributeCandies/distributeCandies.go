package _75DistributeCandies


func distributeCandies(candies []int) int {
	candiesType := make(map[int]int)
	candiesNum := len(candies)
	for _, candy := range candies {
		candiesType[candy]++
	}
	if len(candiesType) > candiesNum/2 {
		return candiesNum/2
	}
	return len(candiesType)
}