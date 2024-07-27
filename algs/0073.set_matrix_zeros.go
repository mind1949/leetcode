package algs

// SetZeroes
//
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
//
// You must do it in place.
//
// Example 1:
//
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
//
// Example 2:
//
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// Constraints:
//
//	m == matrix.length
//	n == matrix[0].length
//	1 <= m, n <= 200
//	-231 <= matrix[i][j] <= 231 - 1
//
// Follow up:
//
//	A straightforward solution using O(mn) space is probably a bad idea.
//	A simple improvement uses O(m + n) space, but still not the best solution.
//	Could you devise a constant space solution?
func SetZeroes(matrix [][]int) {
	// // 直观解法
	// m := len(matrix)
	// n := len(matrix[0])
	// zeros := [2]map[int]bool {
	//     0: make(map[int]bool),
	//     1: make(map[int]bool),
	// }
	// for i := 0; i < m; i++ {
	//     for j := 0; j < n; j++ {
	//         if matrix[i][j] == 0 {
	//             zeros[0][i] = true
	//             zeros[1][j] = true
	//         }
	//     }
	// }

	// for i := 0; i < m; i++ {
	//     for j := 0; j < n; j++ {
	//         if zeros[0][i] || zeros[1][j] {
	//             matrix[i][j] = 0
	//         }
	//     }
	// }

	// 空间复杂度为 O(1) 的解法
	// 使用 matrix 的初行初列代替上述 map 存储待置零行列标记
	//
	//
	// 使用初行初列标记那些行哪些列应该等于0
	var rowZero, columZero bool
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					// 标记第一行中是否有 0
					rowZero = true
				}
				if j == 0 {
					// 标记第一列中是否有0
					columZero = true
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if columZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
