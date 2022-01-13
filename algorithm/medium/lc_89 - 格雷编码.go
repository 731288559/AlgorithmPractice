package medium

import "fmt"

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := 1; i <= n; i++ {
		for j := 1 << (i - 1); j < 1<<i; j++ {
			res[j] = res[1<<i-1-j] ^ 1<<(i-1)
		}
	}
	return res
}

func T_LC89() {
	fmt.Println(grayCode(3))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(1))
}
