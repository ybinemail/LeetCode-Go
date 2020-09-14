package MyContainerWithMostWater

func maxArea(height []int) int {
	maxA := 0
	for i := 0; i < len(height)-1; i++ {
		if height[i] == 0 {
			continue
		}
		for j := i + 1; j < len(height); j++ {
			area := 0
			if height[j] == 0 {
				continue
			}
			if height[i] < height[j] {
				area = height[i] * (j - i)
			} else {
				area = height[j] * (j - i)
			}
			if maxA < area {
				maxA = area
			}
		}
	}
	return maxA
}
