package algs

// Trap solve the following problem:
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
// Example 1:
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
//
// Example 2:
//
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
// Constraints:
//
//	n == height.length
//	1 <= n <= 2 * 104
//	0 <= height[i] <= 105
func Trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var (
		left, right       = 0, len(height) - 1
		leftMax, rightMax int
		water             int
	)
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			water += leftMax - height[left]
			left++
		} else {
			water += rightMax - height[right]
			right--
		}
	}
	return water
}
