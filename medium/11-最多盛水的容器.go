package medium

// 左右边界中间收敛。
func maxArea(height []int) int {
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {

		area := (j - i) * getSmallNum(height[i], height[j])
		if area > maxArea {
			maxArea = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}
	return maxArea
}

func getSmallNum(a, b int) int {
	if a < b {
		return a
	}
	return b

}
