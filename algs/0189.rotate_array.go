package algs

// Rotate
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
// Example 2:
//
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]
//
// Constraints:
//
//	1 <= nums.length <= 10^5
//	-2^31 <= nums[i] <= 2^31 - 1
//	0 <= k <= 10^5
//
// Follow up:
//
//	Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
//	Could you do it in-place with O(1) extra space?
func Rotate(nums []int, k int) {
	// 使用环状替换法
	// 将数组看成环状数组
	n, k := len(nums), k%len(nums)
	cnt := gcd(k, n)
	for i := 0; i < cnt; i++ {
		tmp := nums[i]
		for ok, j := true, (i+k)%n; ok; ok, j = j != i, (j+k)%n {
			nums[j], tmp = tmp, nums[j]
		}
	}
}

// gdc Greatest Common Divisor
// 利用欧几里得算法求计算最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
