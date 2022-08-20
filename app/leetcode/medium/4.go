package medium

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sorted := make([]int, 0, len(nums1)+len(nums2))

	for len(nums1) > 0 && len(nums2) > 0 {
		current1 := nums1[0]
		current2 := nums2[0]

		if current1 > current2 {
			newIndex := sort.SearchInts(nums2, current1)
			sorted = append(sorted, nums2[:newIndex]...)
			nums2 = nums2[newIndex:]
		} else if current2 > current1 {
			newIndex := sort.SearchInts(nums1, current2)
			sorted = append(sorted, nums1[:newIndex]...)
			nums1 = nums1[newIndex:]
		} else {
			sorted = append(sorted, nums1[0])
			nums1 = nums1[1:]
			sorted = append(sorted, nums2[0])
			nums2 = nums2[1:]
		}
	}

	sorted = append(sorted, nums1...)
	sorted = append(sorted, nums2...)

	middle := len(sorted) / 2
	hasEvenLength := len(sorted)%2 == 0
	if !hasEvenLength {
		return float64(sorted[middle])
	}
	return float64(sorted[middle-1]+sorted[middle]) / 2
}

func binarySearch(nums []int, num int) int {
	left, right := 0, len(nums)

	for left <= right {
		switch middle := (right-left)/2 + left; {
		case num < nums[middle]:
			right = middle - 1
		case num > nums[middle]:
			left = middle + 1
		default:
			return middle + 1
		}
	}
	return len(nums)
}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
