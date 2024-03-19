package leetcode

import (
	"sort"
	"strconv"
)

/*
	670. 最大交换

	给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

	示例 1 :

	输入: 2736
	输出: 7236
	解释: 交换数字2和数字7。
	示例 2 :

	输入: 9973
	输出: 9973
	解释: 不需要交换。
	注意:

	给定数字的范围是 [0, 108]
*/

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))

	idx1 := 0
	s2 := make([]byte, len(s))
	copy(s2, s)
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] > s2[j]
	})
	for i := 0; i < len(s); i++ {
		if s[i] != s2[i] {
			idx1 = i
			break
		}
	}

	idx2 := idx1
	maxV := s[idx1]
	for i := idx1; i < len(s); i++ {
		if s[i] >= maxV {
			maxV = s[i]
			idx2 = i
		}
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]

	result, _ := strconv.Atoi(string(s))
	return result
}
