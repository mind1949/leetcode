package algs

// SubarraySum
//
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Example 1:
//
// Input: nums = [1,1,1], k = 2
// Output: 2
//
// Example 2:
//
// Input: nums = [1,2,3], k = 3
// Output: 2
//
// Constraints:
//
//	1 <= nums.length <= 2 * 10^4
//	-1000 <= nums[i] <= 1000
//	-10^7 <= k <= 10^7
func SubarraySum(nums []int, k int) int {
	var res int
	m := map[int]int{
		0: 1,
	}
	for i, sum := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		res += m[sum-k]
		m[sum] += 1
	}
	return res
}
