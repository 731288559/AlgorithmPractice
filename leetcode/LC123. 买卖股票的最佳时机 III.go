package leetcode

import "math"

/*
	给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

	设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit3(prices []int) int {
	const k = 2
	n := len(prices)

	/*
		dp[i][j][k] 含义：前i天，交易j-1次，持仓状态为0/1时的最大收益

		状态转移方程：
		f[i+1][j][0] = max(f[i][j][0], f[i][j-1][1] + prices[i])
		f[i+1][j][1] = max(f[i][j][1], f[i][j][0] - prices[i])

		边界情况：
		1. k<0时，无效
		2. 第一天持仓情况，无效
		3. 第一天未持仓情况，收益初始化为0
	*/

	// 这里i维度初始化为n+1，j维度初始化为k+2，都是为了采用下标控制+初始化的方案，替代if逻辑去做数组越界的判断
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
