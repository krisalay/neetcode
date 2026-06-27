func searchMatrix(matrix [][]int, target int) bool {
    potentialRow := getPotentialRow(matrix, target)
    if potentialRow == -1 {
        return false
    }
    left, right := 0, len(matrix[0])-1
    for left <= right {
        mid := (left + right) / 2
        if matrix[potentialRow][mid] == target {
            return true
        }
        if matrix[potentialRow][mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}

func getPotentialRow(matrix [][]int, target int) int {
    top, bottom := 0, len(matrix)-1
    var mid int
    for top <= bottom {
        mid = (top + bottom) / 2
        if target > matrix[mid][len(matrix[0])-1] {
            top = mid + 1
        } else if target < matrix[mid][0] {
            bottom = mid - 1
        } else {
            return mid
        }
    }
    return -1
}
