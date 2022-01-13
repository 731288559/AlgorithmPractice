package medium

import "math"

func minFlipsMonoIncr2(s string) int {
	p := make([]int, 1, len(s)+1)
	for i, v := range s {
		if v == '1' {
			p = append(p, p[i]+1)
		} else {
			p = append(p, p[i])
		}
	}

	var min = math.MaxInt
	for i := 0; i <= len(s); i++ {
		t := p[i] + (len(s) - i - (p[len(s)] - p[i])) // 左边1的个数+右边0的个数
		if min > t {
			min = t
		}
	}

	return min
}

func minFlipsMonoIncr(s string) int {
	dp1 := make([]int, len(s))
	dp2 := make([]int, len(s))

	for i, c := range s {
		if i == 0 {
			if c == '1' {
				dp1[i]++
			}
			continue
		}
		if c == '0' {
			dp1[i] = dp1[i-1]
			dp2[i] = min(dp1[i-1], dp2[i-1]+1)
		} else {
			dp1[i] = dp1[i-1] + 1
			dp2[i] = min(dp1[i-1], dp2[i-1])
		}
	}

	return dp2[len(s)-1]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func T_LC926() {
	println(minFlipsMonoIncr("00110"), 1)
	println(minFlipsMonoIncr("010110"), 2)
	println(minFlipsMonoIncr("00011000"), 2)
}
