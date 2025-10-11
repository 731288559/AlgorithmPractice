package leetcode

import "slices"

func maximumTotalDamage(power []int) int64 {
	cnt := make(map[int]int)
	for _, p := range power {
		cnt[p]++
	}
	l := make([]int, 0, len(cnt))
	for k := range cnt {
		l = append(l, k)
	}
	slices.Sort(l)

	// 递归解法
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		j := i
		for j > 0 && l[i]-l[j-1] <= 2 {
			j--
		}
		return max(dfs(i-1), dfs(j-1)+l[i]*cnt[l[i]])
	}

	// 递推解法
	f := make([]int, len(l)+1)
	j := 0
	for i, x := range l {
		for l[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(dfs(len(l) - 1))
}

func maximumTotalDamage2(power []int) int64 {
	cnt := make(map[int]int)
	for _, p := range power {
		cnt[p]++
	}
	l := make([]int, 0, len(cnt))
	for k := range cnt {
		l = append(l, k)
	}
	slices.Sort(l)

	// 递推解法（dp）
	f := make([]int, len(l)+1)
	j := 0
	for i, x := range l {
		for l[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[len(l)])
}
