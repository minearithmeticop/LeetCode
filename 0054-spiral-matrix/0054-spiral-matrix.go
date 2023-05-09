func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    m, n := len(matrix), len(matrix[0])
    top, bottom, left, right := 0, m-1, 0, n-1
    result := make([]int, 0, m*n)

    for {
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++
        if top > bottom {
            break
        }

        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--
        if left > right {
            break
        }

        for i := right; i >= left; i-- {
            result = append(result, matrix[bottom][i])
        }
        bottom--
        if top > bottom {
            break
        }

        for i := bottom; i >= top; i-- {
            result = append(result, matrix[i][left])
        }
        left++
        if left > right {
            break
        }
    }

    return result
}