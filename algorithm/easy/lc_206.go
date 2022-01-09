package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tmp = head
	var prev *ListNode

	for tmp != nil {
		next := tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}
	return prev
}

func T_LC206() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	head2 := reverseList(head)
	for head2 != nil {
		println(head2.Val)
		head2 = head2.Next
	}
}
