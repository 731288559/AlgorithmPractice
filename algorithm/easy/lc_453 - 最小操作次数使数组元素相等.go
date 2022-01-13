package easy

import "math"

// 453. 最小操作次数使数组元素相等

/*
	a, b, c --n次-> a', b', c'
	sum(a+b+c) + n*(len-1) = sum(a'+b'+c') = len*a' = len*(min(a,b,c)+n)
	sum + n*(len-1) = len*(min_val+n)
	sum + n*len - n = min_val*len + n*len
	sum - n = min_val*len
	n = sum - min_val*len

	这里最重要的一点是需要理解为什么最小值每次都会参与自增：
	正常模拟的情况下，每轮会将最大值以外的值自增，因此，最小值必然和其他值一起自增，不可能成为新的最大值，所以每次都会参与自增
*/
func minMoves(nums []int) int {
	sum := 0
	min := math.MaxInt64
	for i := range nums {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}

func T_LC453() {
	nums := []int{1, 2, 3}
	println(minMoves(nums))
	nums = []int{1, 1, 1}
	println(minMoves(nums))
	nums = []int{1, 2, 3, 4}
	println(minMoves(nums))
}
