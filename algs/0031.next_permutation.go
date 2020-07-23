package algs

/*
NextPermutation solves the following problem:
	Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

	If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

	The replacement must be in-place and use only constant extra memory.

	Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

		1,2,3 → 1,3,2
		3,2,1 → 1,2,3
		1,1,5 → 1,5,1
*/
func NextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}

	l, r := numsLen-2, numsLen-1
	for l >= 0 && nums[l] >= nums[r] {
		l--
		r--
	}

	if l >= 0 {
		i := numsLen - 1
		for nums[i] <= nums[l] {
			i--
		}
		nums[i], nums[l] = nums[l], nums[i]
	}

	i, j := r, numsLen-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
