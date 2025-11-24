package leetcode

import "fmt"

// 哈希表+递归法
func copyRandomListV1(head *Node) *Node {
	m := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		n, ok := m[node]
		if ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		m[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}

	return deepCopy(head)
}

// 交错链表（相对位置保留映射关系）
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 1. 复制节点，插入到原节点后面
	tmp := head
	for tmp != nil {
		next := tmp.Next
		newNode := &Node{Val: tmp.Val, Next: next}
		tmp.Next = newNode
		tmp = next
	}
	printList(head)

	// 2. 处理random节点关系
	tmp = head
	for tmp != nil {
		if tmp.Random != nil {
			tmp.Next.Random = tmp.Random.Next
		}
		tmp = tmp.Next.Next
	}
	printList(head)

	// 3. 删除原链表节点
	//newHead := head.Next
	//tmp = newHead
	//for tmp.Next != nil && tmp.Next.Next != nil {
	//	next := tmp.Next.Next
	//	tmp.Next = next
	//	tmp = next
	//}
	//return newHead
	// 把交错链表分离成两个链表
	newHead := head.Next
	cur := head
	for ; cur.Next.Next != nil; cur = cur.Next {
		clone := cur.Next
		cur.Next = clone.Next        // 恢复原节点的 next
		clone.Next = clone.Next.Next // 设置新节点的 next
	}
	cur.Next = nil // 恢复原节点的 next
	return newHead
}

func printList(head *Node) {
	tmp := head
	fmt.Println("print list")
	for tmp != nil {
		fmt.Printf("val: %d ", tmp.Val)
		randomVal := 0
		if tmp.Random != nil {
			randomVal = tmp.Random.Val
		}
		fmt.Printf("random val: %d", randomVal)
		tmp = tmp.Next
		fmt.Println()
	}
}
