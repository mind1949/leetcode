package algs

import "sort"

// Merge
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Example 1:
//
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
// Example 2:
//
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Constraints:
//
//	1 <= intervals.length <= 10^4
//	intervals[i].length == 2
//	0 <= starti <= endi <= 10^4
func Merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 合并
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			res = append(res, intervals[0])
			continue
		}
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
			continue
		}
		res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
	}
	return res
}
