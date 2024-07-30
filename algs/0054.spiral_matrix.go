package algs

// SpiralOrder
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
//
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
// Constraints:
//
//	m == matrix.length
//	n == matrix[i].length
//	1 <= m, n <= 10
//	-100 <= matrix[i][j] <= 100
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)

	left, right, top, bottom := 0, n, 0, m
	for left < right && top < bottom && len(res) < cap(res) {
		// 向右
		for i := left; i < right && len(res) < cap(res); i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// 向下
		for i := top; i < bottom && len(res) < cap(res); i++ {
			res = append(res, matrix[i][right-1])
		}
		right--
		// 向左
		for i := right - 1; i >= left && len(res) < cap(res); i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--
		// 向上
		for i := bottom - 1; i > top-1 && len(res) < cap(res); i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
