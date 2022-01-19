package medium

import (
	"container/heap"
	"fmt"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res = make([][]int, 0, k)

	h := make(hp, 0)
	heap.Init(&h)
	for i := 0; i < len(nums1); i++ {
		h.Push([]int{nums1[i] + nums2[0], i, 0})
	}

	for len(h) > 0 && len(res) < k {
		t := heap.Pop(&h).([]int)
		res = append(res, []int{nums1[t[1]], nums2[t[2]]})
		if t[2]+1 < len(nums2) {
			heap.Push(&h, []int{nums1[t[1]] + nums2[t[2]+1], t[1], t[2] + 1})
		}
	}

	return res
}

// 最小堆模板
type hp [][]int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func T_LC373() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{3}, 3))
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{1, 2, 3}, 6))
	fmt.Println(kSmallestPairs([]int{1, 2, 3}, []int{1, 2}, 6))
}
