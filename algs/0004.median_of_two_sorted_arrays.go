package algs

/*
FindMedianSortedArrays sloves the following problem:
	There are two sorted arrays nums1 and nums2 of size m and n respectively.

	Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

	You may assume nums1 and nums2 cannot be both empty.

	Example 1:

		nums1 = [1, 3]
		nums2 = [2]

		The median is 2.0

	Example 2:

		nums1 = [1, 2]
		nums2 = [3, 4]

		The median is (2 + 3)/2 = 2.5
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge
	nums3 := merge(nums1, nums2)
	// find median
	if len(nums3)%2 != 0 {
		return float64(nums3[len(nums3)/2])
	}
	return float64(nums3[len(nums3)/2]+nums3[len(nums3)/2-1]) / 2
}

func merge(nums1, nums2 []int) []int {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		num := 0
		if nums1[i] < nums2[j] {
			num = nums1[i]
			i++
		} else {
			num = nums2[j]
			j++
		}
		nums3 = append(nums3, num)
	}
	if i >= len(nums1) {
		nums3 = append(nums3, nums2[j:]...)
	}
	if j >= len(nums2) {
		nums3 = append(nums3, nums1[i:]...)
	}
	return nums3
}
