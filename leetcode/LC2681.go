package leetcode

import (
	"math"
	"sort"
)

/*
	2681. 英雄的力量

	给你一个下标从 0 开始的整数数组 nums ，它表示英雄的能力值。如果我们选出一部分英雄，这组英雄的 力量 定义为：
	i0 ，i1 ，... ik 表示这组英雄在数组中的下标。那么这组英雄的力量为 max(nums[i0],nums[i1] ... nums[ik])2 * min(nums[i0],nums[i1] ... nums[ik]) 。
	请你返回所有可能的 非空 英雄组的 力量 之和。由于答案可能非常大，请你将结果对 109 + 7 取余。
*/
func sumOfPower(nums []int) int {
	result := 0
	// 1,2,4
	sort.Ints(nums)
	length := len(nums)

	l := int(math.Pow(2, float64(length)))
	for n := 1; n < l; n++ {
		min := 0
		max := 0
		for i := 0; i < length; i++ {
			if n&(1<<i) > 0 {
				min = nums[i]
				break
			}
		}
		for j := length - 1; j >= 0; j-- {
			if n&(1<<j) > 0 {
				max = nums[j]
				break
			}
		}
		curResult := max * max * min
		result = (result + curResult) % (1000000007)
	}
	return result
}

func sumOfPower2(nums []int) (ans int) {
	const mod = 1_000_000_007
	sort.Ints(nums)
	s := 0
	for _, x := range nums { // x 作为最大值
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod               // 递推计算下一个 s
	}
	return
}
