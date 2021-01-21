package main

func main() {

}

/**
二维数组中查找
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	row := len(matrix)
	if row < 1 {
		return false
	}
	col := len(matrix[0]) - 1
	if col < 0 {
		return false
	}
	i := 0
	for i < row && col >= 0 {
		if matrix[i][col] == target {
			return true
		} else if matrix[i][col] > target {
			col--
		} else if matrix[i][col] < target {
			i++
		}
	}
	return false
}
