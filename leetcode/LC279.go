package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minTemp := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minTemp = min(minTemp, dp[i-j*j])
		}
		dp[i] = minTemp + 1
	}

	return dp[n]
}
