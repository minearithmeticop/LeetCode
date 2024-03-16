func removeElement(nums []int, val int) int {
    i := 0 // Pointer for iterating over the array

    // Iterate over the array using pointer i
    for j := 0; j < len(nums); j++ {
        // If the current element is not equal to val, copy it to the i-th position
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }

    // Return the new length of the array
    return i
}
