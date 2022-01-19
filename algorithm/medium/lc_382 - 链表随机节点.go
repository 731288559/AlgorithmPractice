package medium

import (
	"math/rand"
)

// 解法：蓄水池抽样
// https://leetcode-cn.com/problems/linked-list-random-node/solution/xu-shui-chi-chou-yang-suan-fa-by-jackwener/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution382 struct {
	Head *ListNode
}

func Constructor382(head *ListNode) Solution382 {
	return Solution382{Head: head}
}

func (this *Solution382) GetRandom() int {
	var res int
	for tmp, i := this.Head, 1; tmp != nil; i++ {
		if rand.Intn(i) == 0 {
			res = tmp.Val
		}
		tmp = tmp.Next
	}
	return res
}

func T_LC382() {

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	obj := Constructor382(head)
	println(obj.GetRandom())
	println(obj.GetRandom())
	println(obj.GetRandom())
}
