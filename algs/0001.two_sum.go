package algs

/*
TwoSum solves the fllowing problom:

	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

		Given nums = [2, 7, 11, 15], target = 9,

		Because nums[0] + nums[1] = 2 + 7 = 9,
		return [0, 1].
*/
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := numMap[target-v]; ok {
			return []int{j, i}
		}
		numMap[v] = i
	}

	return nil
}
