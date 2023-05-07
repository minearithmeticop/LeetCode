func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    subsequences := make([]int, 0, len(obstacles))
    result := make([]int, 0, len(obstacles))
    for _, obstacle := range obstacles {
        idx := sort.SearchInts(subsequences, obstacle+1)
        if idx == len(subsequences) {
            subsequences = append(subsequences, obstacle)
        } else {
            subsequences[idx] = obstacle
        }
        result = append(result, idx+1)
    }
    return result
}
