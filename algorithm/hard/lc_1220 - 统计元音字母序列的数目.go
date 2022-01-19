package hard

func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7
	dp := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = []int{
			(dp[1] + dp[2] + dp[4]) % mod,
			(dp[0] + dp[2]) % mod,
			(dp[1] + dp[3]) % mod,
			(dp[2]) % mod,
			(dp[2] + dp[3]) % mod,
		}
	}

	sum := 0
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
	}
	return sum % mod
}

func T_LC1220() {
	println(countVowelPermutation(1), 5)
	println(countVowelPermutation(2), 10)
	println(countVowelPermutation(5), 68)
	println(countVowelPermutation(144))
}
