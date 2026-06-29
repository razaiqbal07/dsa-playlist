
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	result := []int{}
	for right > left {
		if numbers[left]+numbers[right] == target {
			result = []int{left + 1, right + 1}
			return result
		}
		if right == left+1 {
			left++
			right = len(numbers)
		}
		right--
	}
	return result
}
