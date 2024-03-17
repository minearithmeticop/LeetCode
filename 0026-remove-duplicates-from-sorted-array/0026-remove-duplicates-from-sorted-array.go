func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    // Initialize pointer i
    i := 1

    // Iterate over the array using pointer j
    for j := 1; j < len(nums); j++ {
        // If nums[j] is different from nums[i-1], copy it to nums[i]
        if nums[j] != nums[i-1] {
            nums[i] = nums[j]
            i++
        }
    }
    
    // Modify nums slice to contain only unique elements
    nums = nums[:i]

    // Return the length of the modified nums slice
    return len(nums)
}
