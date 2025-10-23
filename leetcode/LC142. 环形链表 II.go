package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for slow != fast {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	fake := &ListNode{Next: head}
	slow = fake
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// 变形：如果存在环，进行断尾
func detectCycleV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for slow != fast {
		if slow == nil || fast == nil || fast.Next == nil {
			return head
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	fake := &ListNode{Next: head}
	slow = fake
	for slow != fast {
		if slow.Next == fast.Next {
			fast.Next = nil
			return head
		}
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
