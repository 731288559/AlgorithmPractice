package leetcode

import "math"

/*
	给你一个整数数组 prices，其中 prices[i] 是第 i 天股票的价格（美元），以及一个整数 k。

	你最多可以进行 k 笔交易，每笔交易可以是以下任一类型：

	普通交易：在第 i 天买入，然后在之后的第 j 天卖出，其中 i < j。你的利润是 prices[j] - prices[i]。

	做空交易：在第 i 天卖出，然后在之后的第 j 天买回，其中 i < j。你的利润是 prices[i] - prices[j]。

	注意：你必须在开始下一笔交易之前完成当前交易。此外，你不能在已经进行买入或卖出操作的同一天再次进行买入或卖出操作。

	通过进行 最多 k 笔交易，返回你可以获得的最大总利润。
*/

func maximumProfitV1(prices []int, k int) int64 {
	n := len(prices)

	// 记忆化
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{-1, -1, -1} // -1 表示还没有计算过
		}
	}

	var dfs func(i int, j int, hold int) int
	dfs = func(i int, j int, hold int) (res int) {
		if j < 0 {
			return math.MinInt / 2 // 防止溢出
		}
		if i < 0 {
			if hold != 0 {
				return math.MinInt / 2
			}
			return 0
		}

		ptr := &memo[i][j][hold]
		if *ptr != -1 { // 之前计算过
			return *ptr
		}
		defer func() { *ptr = res }() // 记忆化

		if hold == 1 {
			// 持仓情况下什么也不做 / 前一天空仓买入
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		}

		if hold == 2 {
			// 卖空情况下什么也不做 / 前一天空仓卖空
			return max(dfs(i-1, j, 2), dfs(i-1, j-1, 0)+prices[i])
		}

		// 空仓情况下什么也不做 / 前一天持仓卖出 / 前一天卖空平仓
		return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i], dfs(i-1, j, 2)-prices[i])
	}

	return int64(dfs(n-1, k, 0))
}

// 递推
func maximumProfitV2(prices []int, k int) int64 {
	n := len(prices)

	f := make([][][3]int, n+1)
	for i := range f {
		f[i] = make([][3]int, k+2)
		for j := range f[i] {
			f[i][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		}
	}

	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}

	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p, f[i][j][2]-p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
			f[i+1][j][2] = max(f[i][j][2], f[i][j-1][0]+p)
		}
	}

	return int64(f[n][k+1][0])
}

// 滚动
func maximumProfit(prices []int, k int) int64 {
	//f := make([][][3]int, n+1)
	//for i := range f {
	//	f[i] = make([][3]int, k+2)
	//	for j := range f[i] {
	//		f[i][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	//	}
	//}
	//
	//for j := 1; j <= k+1; j++ {
	//	f[0][j][0] = 0
	//}
	//
	//for i, p := range prices {
	//	for j := 1; j <= k+1; j++ {
	//		f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p, f[i][j][2]-p)
	//		f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
	//		f[i+1][j][2] = max(f[i][j][2], f[i][j-1][0]+p)
	//	}
	//}
	//
	//return int64(f[n][k+1][0])

	// 由于f[i+1]只依赖f[i]，因此最外层循环可以抛弃掉前一次以外的计算结果
	// 由于f[j]依赖f[j-1]，所以反向遍历j
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j] = [3]int{0, math.MinInt / 2, math.MinInt / 2}
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}
