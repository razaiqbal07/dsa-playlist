
func trap(height []int) int {
	trapped := 0
	currentMax := 0
	peak := 0
	left := 0
	right := 0
	accVol := 0
	i := 0
	for i < len(height) {
		if height[i] > height[peak] {
			peak = i
		}
		i++
	}

	for right <= peak {
		if height[right] >= height[left] {
			trapped += accVol
			currentMax = height[right]
			accVol = 0
			left = right
			right++
		} else {
			accVol += currentMax - height[right]
			right++
		}
	}

	currentMax = 0
	accVol = 0
	right = len(height) - 1
	left = len(height) - 1
	for left >= peak {
		if height[left] >= height[right] {
			trapped += accVol
			currentMax = height[left]
			accVol = 0
			right = left
			left--
		} else {
			accVol += currentMax - height[left]
			left--
		}
	}

	return trapped
}