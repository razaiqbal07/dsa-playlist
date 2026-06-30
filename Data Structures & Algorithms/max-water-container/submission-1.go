func maxArea(heights []int) int {
	left :=0
	right := len(heights) - 1
	maximum := math.MinInt
	for left < right {
		h := min(heights[left],heights[right])
		w := right - left
		vol := h * w
		maximum = max(maximum,vol)
		if heights[left] > heights[right] {
			right--
		}else {
			left++
		}
	}
	return maximum
}
