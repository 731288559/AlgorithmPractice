package medium

import "fmt"

// 46. 全排列

func permute(nums []int) [][]int {
	length := 1
	for n := len(nums); n > 0; n-- {
		length *= n
		n--
	}
	result := make([][]int, 0, length)
	visit := make([]bool, len(nums))
	var backtrack func([]int)
	backtrack = func(arr []int) {
		if len(arr) == len(nums) {
			result = append(result, append([]int{}, arr...))
		}
		for i, n := range nums {
			if visit[i] == true {
				continue
			}
			arr = append(arr, n)
			visit[i] = true
			backtrack(arr)
			visit[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	backtrack([]int{})
	return result
}

func T_LC46() {
	// fmt.Println(permute([]int{1, 2, 3}))
	// fmt.Println(permute([]int{0, 1}))
	// fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
