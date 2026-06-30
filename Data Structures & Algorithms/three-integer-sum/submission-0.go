import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		target := -(nums[i])

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}