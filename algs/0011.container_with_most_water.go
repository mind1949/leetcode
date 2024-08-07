package algs

/*
MaxArea solves the following problem:

	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

	Note: You may not slant the container and n is at least 2.

	Example:

		Input: [1,8,6,2,5,4,8,3,7]
		Output: 49
*/
func MaxArea(height []int) int {
	var (
		i, j    = 0, len(height) - 1
		maxArea = 0
	)
	for j-i > 0 {
		area := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, area)
		if height[j] < height[i] {
			j--
		} else {
			i++
		}
	}

	return maxArea
}
