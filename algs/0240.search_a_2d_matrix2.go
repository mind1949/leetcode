package algs

// SearchMatrix
//
// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
//
//	Integers in each row are sorted in ascending from left to right.
//	Integers in each column are sorted in ascending from top to bottom.
//
// Example 1:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// Output: true
//
// Example 2:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// Output: false
//
// Constraints:
//
//	m == matrix.length
//	n == matrix[i].length
//	1 <= n, m <= 300
//	-10^9 <= matrix[i][j] <= 10^9
//	All the integers in each row are sorted in ascending order.
//	All the integers in each column are sorted in ascending order.
//	-10^9 <= target <= 10^9
func SearchMatrix(matrix [][]int, target int) bool {
	// 使用官方题解的 Z字形解法
	if len(matrix) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	x, y := 0, n-1
	for x >= 0 && x < m && y >= 0 && y < n {
		value := matrix[x][y]
		if value == target {
			return true
		} else if value > target {
			y--
		} else {
			x++
		}
	}
	return false
}
