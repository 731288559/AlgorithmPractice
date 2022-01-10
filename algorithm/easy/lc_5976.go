package easy

func checkValid(matrix [][]int) bool {
	var m map[int]struct{}
	for i := 0; i < len(matrix); i++ {
		m = make(map[int]struct{})
		for j := 0; j < len(matrix); j++ {
			if _, ok := m[matrix[i][j]]; ok {
				return false
			} else {
				m[matrix[i][j]] = struct{}{}
			}
		}
	}

	for j := 0; j < len(matrix); j++ {
		m = make(map[int]struct{})
		for i := 0; i < len(matrix); i++ {
			if _, ok := m[matrix[i][j]]; ok {
				return false
			} else {
				m[matrix[i][j]] = struct{}{}
			}
		}
	}

	return true
}

func T_LC5976() {
	m := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	println(checkValid(m), true)
	m2 := [][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}
	println(checkValid(m2), false)
}
