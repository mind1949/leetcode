package algs

import (
	"sort"
)

/*
ThreeSum solves the following problem:
	Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

	Note:

	The solution set must not contain duplicate triplets.

	Example:

		Given array nums = [-1, 0, 1, 2, -1, -4],

		A solution set is:
		[
			[-1, 0, 1],
			[-1, -1, 2]
		]
*/
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 升序排序
	sort.IntSlice(nums).Sort()
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}

	zeros := make([][]int, 0)
	for m := 1; m <= len(nums)-2; m++ {
		l, r := 0, len(nums)-1
		// 中间去重
		if m-1 >= 0 && nums[m] == nums[m-1] {
			if m-2 >= 0 && nums[m-2] == nums[m-1] {
				continue
			}
			l, r = m-1, len(nums)-1
		}
		for l < m && r > m {
			sum := nums[l] + nums[m] + nums[r]
			if sum == 0 {
				zero := []int{nums[l], nums[m], nums[r]}
				zeros = append(zeros, zero)
			}
			if sum > 0 {
				// 右边去重
				for cur := r; r > m && nums[cur] == nums[r]; cur, r = r, r-1 {
				}
			} else {
				// 左边去重
				for cur := l; l < m && nums[cur] == nums[l]; cur, l = l, l+1 {
				}
			}
		}
	}

	return zeros
}
