package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeNode := &ListNode{Next: head}
	partPre := fakeNode
	partHead := head

	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		next := cur.Next
		if cnt == k {
			newHead, newTail := reversePart(partHead, next)
			partPre.Next = newHead
			newTail.Next = next
			partPre = newTail
			partHead = next
			cnt = 0
		}

		cur = next
	}

	return fakeNode.Next
}

func reversePart(head, next *ListNode) (newHead, newTail *ListNode) {
	var pre *ListNode
	cur := head
	for cur != next {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre, head
}
