package other

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteListMN(head *ListNode, m, n int) {
	if m < n {
		m, n = n, m
	}

	dummy := &ListNode{Next: head}
	var n1 *ListNode = dummy
	var n2 *ListNode = dummy
	for i := 0; i < m; i++ {
		if i < n {
			n1 = n1.Next
		}
		n2 = n2.Next
	}

	tmp := dummy
	// TODO
	var prev *ListNode

	var f2 bool
	for n1.Next != nil {
		if !f2 && n2.Next == nil {
			prev.Next = tmp.Next
			f2 = true
		} else {
			n2 = n2.Next
		}
		prev = tmp
		tmp = tmp.Next
		n1 = n1.Next
	}
	prev.Next = tmp.Next
}

func TestDeleteListMN() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{Val: 6},
					}}},
		},
	}
	deleteListMN(root, 3, 2)
	for root != nil {
		println(root.Val)
		root = root.Next
	}
}
