package medium

// 240. 搜索二维矩阵 II

func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0

	for row >= 0 && col < len(matrix[0]) {
		if target > matrix[row][col] {
			col++
		} else if target < matrix[row][col] {
			row--
		} else {
			return true
		}
	}
	return false
}

func T_LC240() {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	println(searchMatrix2(m, 5))
	println(searchMatrix2(m, 20))
}
