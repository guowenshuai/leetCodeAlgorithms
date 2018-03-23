package _61ArrayPartitionI

//import "sort"

func arrayPairSum(nums []int) int {
	sum := 0
	qSort(nums)
	//sort.Ints(nums)
	for i, n := range nums {
		if i%2 == 0 {
			sum += n
		}
	}
	return sum
}

func qSort(nums []int) {
	var recurese func(left int, right int)

	recurese = func(left int, right int) {
		if left >= right {
			return
		}
		i, j := left, right
		i++
		for i < j {
			if nums[j] < nums[left] {
				nums[i], nums[j] = nums[j], nums[i]
				i++;
			} else {
				j--
			}
		}
		if nums[left] > nums[i] {
			nums[left], nums[i] = nums[i], nums[left]
		}

		recurese(0, i-1)
		recurese(i, right)
	}

	recurese(0, len(nums)-1)
}