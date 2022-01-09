package medium

import (
	"fmt"
	"sort"
)

// 5186. 区间内查询数字的频率

type RangeFreqQuery struct {
	m map[int]sort.IntSlice
}

func Constructor_(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{
		m: make(map[int]sort.IntSlice),
	}
	for i, n := range arr {
		r.m[n] = append(r.m[n], i)
	}
	return r
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	ans := 0
	if l, ok := this.m[value]; ok {
		// fmt.Println(l.Search(left))
		// fmt.Println(l[l.Search(left):])
		// fmt.Println(l[l.Search(left):].Search(right + 1))
		ans = l[l.Search(left):].Search(right + 1)
	}
	return ans
}

func T_LC5186() {
	obj := Constructor_([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
	println(obj.Query(1, 2, 4), 1)
	println(obj.Query(0, 11, 33), 2)

	l := sort.IntSlice{0, 2, 7, 9}
	fmt.Println(l.Search(1))
	fmt.Println(l.Search(8))
	left, right := 1, 8
	fmt.Println(l[l.Search(left):].Search(right + 1))

}
