package algs

// LongestConsecutive solves the following problem:
//
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
// Example 1:
//
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
//
// Example 2:
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
//
// Constraints:
//
//	0 <= nums.length <= 10^5
//	-10^9 <= nums[i] <= 10^9
func LongestConsecutive(nums []int) int {
	exist := make(map[int]bool)
	for _, num := range nums {
		exist[num] = true
	}

	var longest int
	for num := range exist {
		if exist[num-1] {
			continue
		}

		length := 1
		for n := num + 1; exist[n]; n++ {
			length++
		}
		longest = max(longest, length)
	}
	return longest
}
