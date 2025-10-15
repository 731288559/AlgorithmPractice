package leetcode

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(i, j int, board [][]byte)
	dfs = func(i, j int, board [][]byte) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if board[i][j] == 'O' {
			board[i][j] = 'K'
		} else {
			return
		}
		dfs(i, j+1, board)
		dfs(i, j-1, board)
		dfs(i+1, j, board)
		dfs(i-1, j, board)
	}

	for i := range m {
		for j := range n {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				dfs(i, j, board)
			}
		}
	}

	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'K' {
				board[i][j] = 'O'
			}
		}
	}

	return
}
