package leetcode

/*
	84. 柱状图中最大的矩形

	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
	求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}

	var stack []int
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	var ans int
	for i, h := range heights {
		ans = max(ans, (right[i]-left[i]-1)*h)
	}

	return ans
}
