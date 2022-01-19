package other

import "fmt"

func eightQueen() [][][]int {
	var res [][][]int
	n := 4

	var chess = make([][]int, n)
	for i := range chess {
		chess[i] = make([]int, n)
	}

	var helper func(i int)
	var check func(i, j int) bool

	helper = func(i int) {
		if i >= n {
			tmp := [][]int{}
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if chess[i][j] == 1 {
						tmp = append(tmp, []int{i, j})
					}
				}
			}
			res = append(res, tmp)
			return
		}

		for j := 0; j < n; j++ {
			if check(i, j) {
				chess[i][j] = 1
				helper(i + 1)
				chess[i][j] = 0
			}
		}
	}

	check = func(row, column int) bool {
		for i := 0; i < n; i++ {
			if chess[i][column] == 1 {
				return false
			}
		}
		for i := 0; i <= row+column; i++ {
			j := row + column - i
			if i >= n || j < 0 || j >= n {
				continue
			}
			if chess[i][j] == 1 {
				return false
			}
		}
		for i := 0; i < n; i++ {
			j := column - row + i
			if j < 0 || j >= n {
				continue
			}
			if chess[i][j] == 1 {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		chess[0][i] = 1
		helper(1)
		chess[0][i] = 0
	}

	return res
}

func TestEightQueens() {
	fmt.Println(eightQueen())
}
