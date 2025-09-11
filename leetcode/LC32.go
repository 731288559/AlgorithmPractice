package leetcode

// 32. 最长有效括号

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n+1) // 以s[i]结尾的字符串的子串长度

	for i := 1; i < n; i++ {
		if s[i] == '(' { // 无效
			dp[i] = 0
		}
		if s[i] == ')' {
			if s[i-1] == '(' { // ...() 类型
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

			}
			if s[i-1] == ')' { // ...)) 类型
				if dp[i-1] == 0 {
					dp[i] = 0
					continue
				}

				if i-dp[i-1]-1 >= 0 {
					if s[i-dp[i-1]-1] == '(' {
						/*
								( ) ( ( ) )
							    0 2 0 0 2 ?
							1. 求dp[5] = ?
							2. 发现s[5-dp[4]-1] == '('，正好凑成一对
							3. 计算新对左侧，新对中间和新对的长度和：dp[4] + dp[1] + 2
						*/
						if i-dp[i-1]-2 >= 0 {
							dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}
