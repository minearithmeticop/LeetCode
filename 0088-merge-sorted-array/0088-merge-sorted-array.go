func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1, p2 := m-1, n-1 // Initialize pointers to the end of nums1 and nums2
	p := m + n - 1     // Initialize pointer to the end of the merged array

	// Loop until both nums1 and nums2 still have elements
	for p1 >= 0 && p2 >= 0 {
		// Compare the elements pointed to by p1 and p2
		// If the element at p1 is greater, move it to the end of the merged array
		// Otherwise, move the element at p2 to the end of the merged array
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1-- // Move pointer p1 to the left
		} else {
			nums1[p] = nums2[p2]
			p2-- // Move pointer p2 to the left
		}
		p-- // Move the merged array pointer to the left
	}

	// After the loop, if there are remaining elements in nums2, copy them to nums1
	// Since nums1 has enough space to accommodate nums2, we only need to consider nums2 here
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2-- // Move pointer p2 to the left
		p--  // Move the merged array pointer to the left
	}

	return nums1
}
