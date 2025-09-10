package leetcode

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007

	ans := 0

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i <= n; i++ {
		dp[i] %= mod
		if i+forget-1 >= n {
			ans += dp[i]
		}
		for j := i + delay; j <= min(i+forget-1, n); j++ {
			dp[j] += dp[i]
		}
	}

	return ans % mod
}
