package medium

import "fmt"

// 260. 只出现一次的数字 III

func singleNumber(nums []int) []int {
	result := make([]int, 2)

	x := 0
	for _, i := range nums {
		x ^= i
	}
	bitIdx := 1
	for {
		if x&bitIdx > 0 {
			break
		}
		bitIdx <<= 1
	}
	for _, i := range nums {
		if i&bitIdx > 0 {
			result[0] ^= i
		} else {
			result[1] ^= i
		}
	}
	return result
}

func T_LC260() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
