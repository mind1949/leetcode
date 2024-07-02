package algs

import "math"

// MiniPathSum solves the following problem:
//
//	Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
//
//	Note: You can only move either down or right at any point in time.
//
//
//
//	Example 1:
//
//	1  3  1
//	1  5  1
//	4  2  1
//
//	Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
//	Output: 7
//	Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
//
//	Example 2:
//
//	Input: grid = [[1,2,3],[4,5,6]]
//	Output: 12
//
//	Constraints:
//
//	    m == grid.length
//	    n == grid[i].length
//	    1 <= m, n <= 200
//	    0 <= grid[i][j] <= 200
func MiniPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return math.MaxInt
	}
	m := len(grid[0])
	if m == 0 {
		return math.MaxInt
	}

	return MiniPathSumWithDP(grid)
}

// MiniPathSumWithDFS 使用暴力搜索计算最短路径
//
// 状态转移方程:
// * f(i, j) = min(f(i-1, j), f(i, j-1)) + grid[i][j]
//
// 初始状态:
// * f(0, 0) = grid[0][0]
//
// 边界条件:
// * f(i, j) 中的 i 与 j 均大于等于0
func MiniPathSumWithDFS(grid [][]int, i, j int) int {
	// 初始状态
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	// 不合法状态
	if i < 0 || j < 0 {
		return math.MaxInt
	}

	return min(MiniPathSumWithDFS(grid, i-1, j),
		MiniPathSumWithDFS(grid, i, j-1)) + grid[i][j]
}

// MiniPathSumWithDFSMem 使用记忆搜索优化后的暴力搜索计算最短路径
//
// 状态转移方程:
// * f(i, j) = min(f(i-1, j), f(i, j-1)) + grid[i][j]
//
// 初始状态:
// * f(0, 0) = grid[0][0]
//
// 边界条件:
// * f(i, j) 中的 i 与 j 均大于等于0
func MiniPathSumWithDFSMem(mem, grid [][]int, i, j int) int {
	// 初始状态
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	// 不合法状态
	if i < 0 || j < 0 {
		return math.MaxInt
	}

	if result := mem[i][j]; result != -1 {
		return result
	}
	up := MiniPathSumWithDFS(grid, i-1, j)
	left := MiniPathSumWithDFS(grid, i, j-1)
	mem[i][j] = min(up, left) + grid[i][j]
	return mem[i][j]
}

// MiniPathSumWithDP 使用动态规划解决
//
// 状态转移方程:
// * f(i, j) = min(f(i-1, j), f(i, j-1)) + grid[i][j]
//
// 初始状态:
// * f(0, 0) = grid[0][0]
//
// 边界条件:
// * f(i, j) 中的 i 与 j 均大于等于0
func MiniPathSumWithDP(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// 初始化 dp 表
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 准备初始状态
	dp[0][0] = grid[0][0]
	// 首行状态转移
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	// 首列状态转移
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 逐个遍历
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			// 状态转移方程
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

// NewMiniPathSumBacktrack 使用回溯算法枚举计算最短路径
func NewMiniPathSumBacktrack(grid [][]int) *MiniPathSumBacktrack {
	return &MiniPathSumBacktrack{
		grid: grid,
		high: len(grid),
		wide: len(grid[0]),

		result: math.MaxInt,
	}
}

// MiniPathSumBacktrack 使用回溯算法枚举计算最短路径
type MiniPathSumBacktrack struct {
	// 要走的网格
	grid [][]int
	// 网格的高于宽
	high, wide int

	// 当前所处坐标
	coordinate [2]int
	// 最近位移
	latest [2]int
	// 最近位移路径和
	sum int

	// 结果
	result int
}

func (m *MiniPathSumBacktrack) Calc() int {
	m.sum = m.grid[0][0]
	m.backtrack()
	return m.result
}

func (m *MiniPathSumBacktrack) backtrack() {
	// 判断是否为解
	if m.isSolution() {
		// 记录解
		m.recordSolution()
		return
	}

	// 遍历选择
	for _, choice := range [][2]int{{0, 1}, {1, 0}} {
		// 剪枝
		if m.isValid(choice) {
			pre := m.latest
			// 试探
			m.makeChoice(choice)
			// 回溯
			m.backtrack()
			// 回退
			m.undoChoice(pre)
		}
	}
}

func (m *MiniPathSumBacktrack) isSolution() bool {
	// 是否到达右下角
	return m.coordinate[0] == m.high-1 && m.coordinate[1] == m.wide-1
}

func (m *MiniPathSumBacktrack) recordSolution() {
	if m.sum < m.result {
		m.result = m.sum
	}
}

func (m *MiniPathSumBacktrack) isValid(choice [2]int) bool {
	// 没有越过网格高与宽
	if m.coordinate[0]+choice[0] > m.high-1 {
		return false
	}
	if m.coordinate[1]+choice[1] > m.wide-1 {
		return false
	}
	return true
}

func (m *MiniPathSumBacktrack) makeChoice(choice [2]int) {
	// 移动坐标
	m.coordinate[0] += choice[0]
	m.coordinate[1] += choice[1]
	// 计算路径和
	m.sum += m.grid[m.coordinate[0]][m.coordinate[1]]
	// 更新最近位移坐标
	m.latest = choice
}

func (m *MiniPathSumBacktrack) undoChoice(pre [2]int) {
	// 缩减路径和
	m.sum -= m.grid[m.coordinate[0]][m.coordinate[1]]
	// 回退坐标
	m.coordinate[0] -= m.latest[0]
	m.coordinate[1] -= m.latest[1]
	// 更新最近位移坐标
	m.latest = pre
}
