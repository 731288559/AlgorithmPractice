package medium

import (
	"fmt"
	"sort"
)

// 47. 全排列 II

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
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
			if i > 0 && nums[i-1] == n && !visit[i-1] {
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

func T_LC47() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
