package medium

import (
	"fmt"
)

// 78. 子集

// 方法1: 01标记，遍历2的n次方种组合
// 方法2: 从[[]]开始，遍历当前集合，每次增加一个数字到集合的每个元素中
// 方法3: dfs 回溯法

func subsets(nums []int) [][]int {
	result := make([][]int, 1, 1<<len(nums))
	result[0] = []int{}
	for _, n := range nums {
		for _, arr := range result {
			tmp := make([]int, len(arr), len(arr)+1)
			copy(tmp, arr)
			tmp = append(tmp, n)
			result = append(result, tmp)
		}
	}
	return result
}

func subsets_v3(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func T_LC78() {
	// fmt.Println(subsets([]int{1, 2, 3}))
	// fmt.Println(subsets([]int{0}))
	// fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
