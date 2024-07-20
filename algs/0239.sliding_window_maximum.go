package algs

// MaxSlidingWindow
//
// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
//
// Return the max sliding window.
//
// Example 1:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
//
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]
//
// Constraints:
//
//	1 <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4
//	1 <= k <= nums.length
func MaxSlidingWindow(nums []int, k int) []int {
	// 使用单调队列法
	topIdx := []int{}
	push := func(i int) {
		for len(topIdx) > 0 && nums[i] >= nums[topIdx[len(topIdx)-1]] {
			topIdx = topIdx[:len(topIdx)-1]
		}
		topIdx = append(topIdx, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	res := make([]int, 0, len(nums)-k+1)
	res = append(res, nums[topIdx[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		for topIdx[0] <= i-k {
			topIdx = topIdx[1:]
		}
		res = append(res, nums[topIdx[0]])
	}
	return res
}
