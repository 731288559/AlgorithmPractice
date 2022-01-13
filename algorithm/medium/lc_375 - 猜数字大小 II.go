package medium

import (
	"math"
)

// 375. 猜数字大小 II

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, n+1)
	}

	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			if j == i+1 {
				dp[i][j] = i
				continue
			}
			if j == i+2 {
				dp[i][j] = i + 1
				continue
			}
			minCost := math.MaxInt
			for k := i + 1; k <= j-1; k++ {
				cost := k + maxInt(dp[i][k-1], dp[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			dp[i][j] = minCost
		}
	}

	// for i := range dp {
	// 	fmt.Println(dp[i])
	// }
	return dp[1][n]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func T_LC375() {
	println(getMoneyAmount(5), 6)

	println(getMoneyAmount(10), 16)
	println(getMoneyAmount(1), 0)
	println(getMoneyAmount(2), 1)
}
