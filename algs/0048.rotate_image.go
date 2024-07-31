package algs

// RotateMatrix
//
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
//
// Example 2:
//
// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
// Constraints:
//
//	n == matrix.length == matrix[i].length
//	1 <= n <= 20
//	-1000 <= matrix[i][j] <= 1000
func RotateImage(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	n := len(matrix)
	// 原地旋转九十度
	// 可以看做是从矩阵最外层旋转至最内层
	// 通过观察可以得知，元素旋转规则为:
	// * 列索引 l 变成行索引 r
	// * 行索引 r 通过公式 n -1 -r 变成列索引 l
	for i := 0; i < n/2+n%2; i++ {
		for j := n - 1 - i; j > i; j-- {
			r, l := i, n-1-j
			tmp := matrix[r][l]
			matrix[r][l] = matrix[j][i]
			for {
				r, l = l, n-1-r
				matrix[r][l], tmp = tmp, matrix[r][l]
				if r == j && l == i {
					break
				}
			}
		}
	}
}
