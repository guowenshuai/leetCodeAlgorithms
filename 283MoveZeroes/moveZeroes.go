package _83MoveZeroes

func moveZeroes(nums []int)  {

	for i,j,count := 0,len(nums),0; i<j; i++ {
		if nums[i] == 0 {
			count++
		} else {
			if(count > 0) {
				nums[i-count], nums[i] = nums[i], 0
			}
		}
	}
}
