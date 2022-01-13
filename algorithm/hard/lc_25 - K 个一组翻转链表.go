package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var root = &ListNode{Next: head}
	var pre = root
	var tmp = root

	for tmp != nil {
		for i := 0; i < k; i++ {
			tmp = tmp.Next
			if tmp == nil {
				return root.Next
			}
		}
		next := tmp.Next
		tmp.Next = nil
		newHead, newTail := myReverse(pre.Next)
		newTail.Next = next
		pre.Next = newHead
		pre = newTail
		tmp = newTail
	}
	return root.Next
}

func myReverse(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev, head
}

func MakeList(l []int) *ListNode {
	var head *ListNode = &ListNode{}
	var cur = head
	var tmp *ListNode
	for _, v := range l {
		tmp = &ListNode{Val: v}
		cur.Next = tmp
		cur = cur.Next
	}
	return head.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		print(head.Val)
		head = head.Next
	}
	println()
}

func T_LC25() {
	head := MakeList([]int{1, 2, 3, 4, 5})
	PrintList(head)

	head2 := reverseKGroup(head, 4)
	PrintList(head2)
}
