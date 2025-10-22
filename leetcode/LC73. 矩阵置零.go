package leetcode

// 解法一
// 时间复杂度：O(mn)
// 空间复杂度：O(m + n)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := range m {
		for j := range n {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 解法二
// 时间复杂度：O(mn)
// 空间复杂度：O(1)
func setZeroesV2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	row0, col0 := false, false
	for i := range m {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}
	for j := range n {
		if matrix[0][j] == 0 {
			row0 = true
			break
		}
	}

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if col0 {
		for i := range m {
			matrix[i][0] = 0
		}
	}

	if row0 {
		for j := range n {
			matrix[0][j] = 0
		}
	}
}
