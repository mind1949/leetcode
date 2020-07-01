package algs

import (
	"math"
	"sort"
)

/*
ThreeSumClosest solves the following problem:
	Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

	Example 1:

		Input: nums = [-1,2,1,-4], target = 1
		Output: 2
		Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).



	Constraints:

		3 <= nums.length <= 10^3
		-10^3 <= nums[i] <= 10^3
		-10^4 <= target <= 10^4
*/
func ThreeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)

	// 查找里target最近的组合
	var closest = int(math.MaxInt64)
	for l := 0; l < len(nums)-2; l++ { // left
		if l > 0 && nums[l] == nums[l-1] { // left 去重
			continue
		}
		var (
			m = l + 1
			r = len(nums) - 1
		)
		for m < r || closest == target {
			sum := nums[r] + nums[m] + nums[l]
			if math.Abs(float64(sum-target)) <
				math.Abs(float64(closest-target)) { // 判断谁离target更近
				closest = sum
			}

			distance := sum - target
			if distance > 0 {
				for cur := r; r > m && nums[cur] == nums[r]; r-- { // right 去重
				}
			}
			if distance < 0 {
				for cur := m; m < r && nums[cur] == nums[m]; m++ { // middle 去重
				}
			}
			if distance == 0 {
				return sum
			}
		}
	}

	return closest
}
