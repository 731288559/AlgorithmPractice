package leetcode

func canReach(arr []int, start int) bool {
	n := len(arr)
	mark := make([]bool, n)

	var dfs func(arr []int, mark []bool, k int) bool

	dfs = func(arr []int, mark []bool, k int) bool {
		// 数组越界
		if k < 0 || k >= n {
			return false
		}

		if arr[k] == 0 {
			return true
		}

		// 已经走过
		if mark[k] {
			return false
		}
		mark[k] = true
		return dfs(arr, mark, k+arr[k]) || dfs(arr, mark, k-arr[k])
	}

	return dfs(arr, mark, start)
}
