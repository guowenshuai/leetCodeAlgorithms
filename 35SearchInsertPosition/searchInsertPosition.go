package _5SearchInsertPosition


func searchInsert(nums []int, target int) int {
	lenNums := len(nums)

	if lenNums == 0 {
		return 0
	}
	if lenNums > 0 && nums[lenNums]<target {
		return lenNums
	}
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return 0
}
