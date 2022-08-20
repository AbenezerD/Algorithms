package easy

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := 0, 0, m
	swapped := false
	for ; i < k && j < n; i++ {
		if swapped {
			if nums1[k] < nums1[i] {
				nums1[k], nums1[i] = nums1[i], nums1[k]
				swapped = false
				k++
			}
		}
		if nums2[j] < nums1[i] {
			nums1[k] = nums1[i]
			nums1[i] = nums2[j]
			j++
			k++
			swapped = true
		}
	}

	if j != 0 {
		i++
	}
	for ; j < n; j++ {
		nums1[i] = nums2[j]
		i++
	}

	return nums1
}

func _merge(nums1 []int, m int, nums2 []int, n int) []int {
	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	i, j := 0, m
	for i < m+n {
		if nums1[j] < nums1[i] {
			nums1[i], nums1[j] = nums1[j], nums1[i]
		}
		if i == j {
			j++
		} else {
			i++
		}
	}

	return nums1
}
