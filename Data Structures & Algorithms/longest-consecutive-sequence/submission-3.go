func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maximum := 1
	currentMax := 1
	var left int
	var right int
	for i := 0; i < len(nums); i++ {
		left = nums[i] - 1
		right = nums[i] + 1
		for j := 0; j < len(nums); j++ {
			if nums[j] == right {
				currentMax = currentMax + 1
				right = right + 1
				j = 0
			}
			if nums[j] == left {
				currentMax = currentMax + 1
				left = left - 1
				j = 0
			}
		}
		maximum = max(maximum, currentMax)
		currentMax = 1
	}
	return maximum
}
