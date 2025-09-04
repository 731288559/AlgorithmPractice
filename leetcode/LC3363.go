package leetcode

import "math"

/*
	3363. 最多可收集的水果数目

	有一个游戏，游戏由 n x n 个房间网格状排布组成。

	给你一个大小为 n x n 的二维整数数组 fruits ，其中 fruits[i][j] 表示房间 (i, j) 中的水果数目。有三个小朋友 一开始 分别从角落房间 (0, 0) ，(0, n - 1) 和 (n - 1, 0) 出发。

	每一位小朋友都会 恰好 移动 n - 1 次，并到达房间 (n - 1, n - 1) ：

	从 (0, 0) 出发的小朋友每次移动从房间 (i, j) 出发，可以到达 (i + 1, j + 1) ，(i + 1, j) 和 (i, j + 1) 房间之一（如果存在）。
	从 (0, n - 1) 出发的小朋友每次移动从房间 (i, j) 出发，可以到达房间 (i + 1, j - 1) ，(i + 1, j) 和 (i + 1, j + 1) 房间之一（如果存在）。
	从 (n - 1, 0) 出发的小朋友每次移动从房间 (i, j) 出发，可以到达房间 (i - 1, j + 1) ，(i, j + 1) 和 (i + 1, j + 1) 房间之一（如果存在）。
	当一个小朋友到达一个房间时，会把这个房间里所有的水果都收集起来。如果有两个或者更多小朋友进入同一个房间，只有一个小朋友能收集这个房间的水果。当小朋友离开一个房间时，这个房间里不会再有水果。

	请你返回三个小朋友总共 最多 可以收集多少个水果。
*/

func maxCollectedFruits(fruits [][]int) int {
	res := 0

	// 第一个小朋友只能走对角线，否则无法在n-1步后到达终点
	n := len(fruits)
	for i := 0; i < n; i++ {
		res += fruits[i][i]
		fruits[i][i] = math.MinInt
	}

	// 第二个、第三个小朋友只能在各自的三角区域行动，否则也无法走到终点
	// 由于对角线必定被第一个人走过，可以排除掉，因此只需要在(n-1)*(n-1)的三角区域内活动

	dp2, dp3 := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		dp2[i] = make([]int, n)
		dp3[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp2[i][j] = math.MinInt
			dp3[i][j] = math.MinInt
		}
	}

	// 从右上角出发
	dp2[0][n-1] = fruits[0][n-1]
	dx := [3]int{-1, -1, -1}
	dy := [3]int{-1, 0, 1}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < 3; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x < 0 || x >= n || y < 0 || y >= n {
					continue
				}
				if x >= y {
					continue
				}
				tmp := fruits[i][j] + dp2[x][y]
				dp2[i][j] = max(dp2[i][j], tmp)
			}
		}
	}
	// 从左下角出发
	dp3[n-1][0] = fruits[n-1][0]
	dx = [3]int{-1, 0, 1}
	dy = [3]int{-1, -1, -1}
	for j := 1; j < n; j++ {
		for i := j + 1; i < n; i++ {
			for k := 0; k < 3; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x < 0 || x >= n || y < 0 || y >= n {
					continue
				}
				if x <= y {
					continue
				}
				tmp := fruits[i][j] + dp3[x][y]
				dp3[i][j] = max(dp3[i][j], tmp)
			}
		}
	}

	return res + dp2[n-2][n-1] + dp3[n-1][n-2]
}
