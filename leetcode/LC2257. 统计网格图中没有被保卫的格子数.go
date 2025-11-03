package leetcode

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]string, m)
	for i := range m {
		grid[i] = make([]string, n)
	}
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = "wall"
	}
	for _, guard := range guards {
		grid[guard[0]][guard[1]] = "guard"
	}

	ans := 0

	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	markGuard := func(row, col int) {
		for _, dir := range dirs {
			x, y := row+dir[0], col+dir[1]
			stop := false
			for x >= 0 && x < m && y >= 0 && y < n && !stop {
				switch grid[x][y] {
				case "wall", "guard":
					stop = true
				case "":
					ans++
					grid[x][y] = "see"
				}
				x += dir[0]
				y += dir[1]
			}
		}
	}

	for _, guard := range guards {
		markGuard(guard[0], guard[1])
	}

	return m*n - len(guards) - len(walls) - ans
}
