package medium

// 74. 搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0

	for col >= 0 && col < len(matrix[0]) && row >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

func T_LC74() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	println(searchMatrix(m, 3))
	println(searchMatrix(m, 13))
}
