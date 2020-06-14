package algs

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cs := []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: []int{1, 2, 3},
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: []int{1, 2, 3, 4},
		},
	}
	for _, c := range cs {
		got := merge(c.nums1, c.nums2)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("nums1: %v, nums2, %v, expect: %v, got: %v", c.nums1, c.nums2, c.expect, got)
		}
	}

}

func TestFindMedianSortedArrays(t *testing.T) {
	cs := []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2.0,
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: 2.5,
		},
	}

	for _, c := range cs {
		got := FindMedianSortedArrays(c.nums1, c.nums2)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("nums1: %v, nums2, %v, expect: %f, got: %f", c.nums1, c.nums2, c.expect, got)
		}
	}
}
