func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // Initialize pointer i
    i := 1

    // Iterate over the array using pointer j
    for j := 1; j < len(nums); j++ {
        // If nums[j] is different from nums[i-1], copy it to nums[i] and increment i
        if nums[j] != nums[i-1] {
            nums[i] = nums[j]
            i++
        }
    }

    // Return the new length of the array
    return i
}
