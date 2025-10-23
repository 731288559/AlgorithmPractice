package leetcode

import (
	"fmt"
	"testing"
)

func Test_detectCycleV2(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n1

	ans := detectCycleV2(n1)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}
