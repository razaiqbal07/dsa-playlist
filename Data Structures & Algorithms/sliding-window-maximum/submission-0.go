import (
	"slices"
)

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := k
	result := []int{}
	for right <= len(nums) {
		slice := nums[left:right]
		max := slices.Max(slice)
		result = append(result, max)
		right++
		left++
	}
	return result
}