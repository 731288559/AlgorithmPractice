package leetcode

import "math"

/*
	给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。

	设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。

	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

// DFS版本
func maxProfit4V1(k int, prices []int) int {
	n := len(prices)

	var dfs func(i int, j int, hold bool) int
	dfs = func(i int, j int, hold bool) int {
		if j < 0 {
			return math.MinInt
		}
		if i < 0 {
			if hold {
				return math.MinInt
			}
			return 0
		}
		if hold {
			// 持有情况下什么也不做 / 前一天买入
			return max(dfs(i-1, j, true), dfs(i-1, j, false)-prices[i])
		}
		// 未持有情况下什么也不做 / 前一天卖出
		return max(dfs(i-1, j, false), dfs(i-1, j-1, true)+prices[i])
	}

	return dfs(n-1, k, false)
}

// 递推版本
func maxProfit4(k int, prices []int) int {
	n := len(prices)

	/*
		状态转移方程：
		f[i+1][j][0] = max(f[i][j][0], f[i][j-1][1] + prices[i])
		f[i+1][j][1] = max(f[i][j][1], f[i][j][0] - prices[i])

		边界情况：
		1. k<0时，无效
		2. 第一天持仓情况，无效
		3. 第一天未持仓情况，收益初始化为0
	*/
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, k+2)
		for j := range dp[i] {
			dp[i][j] = [2]int{math.MinInt, math.MinInt}
		}
	}
	for j := 1; j <= k+1; j++ {
		dp[0][j][0] = 0
	}

	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			dp[i+1][j][0] = max(dp[i][j][0], dp[i][j-1][1]+p)
			dp[i+1][j][1] = max(dp[i][j][1], dp[i][j][0]-p)
		}
	}

	return dp[n][k+1][0]
}
