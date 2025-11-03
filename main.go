package main

import "fmt"

func main() {

	// m个协程，完成n个任务，等待所有任务完成后退出
	// 不使用 waitGroup

	f()
	runtime.Timer
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树转先序链表
func f() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	fakeNode := &ListNode{}
	cur := fakeNode

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		fmt.Println(root.Val)
		cur.Next = &ListNode{Val: root.Val}
		cur = cur.Next

		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	printList(fakeNode.Next)

	fakeNode = &ListNode{}
	cur = fakeNode

	var midOrder func(root *TreeNode)
	midOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		midOrder(root.Left)
		fmt.Println(root.Val)
		cur.Next = &ListNode{Val: root.Val}
		cur = cur.Next
		midOrder(root.Right)
	}

	midOrder(root)
	printList(fakeNode.Next)
}

func printList(root *ListNode) {
	tmp := root
	for tmp != nil {
		fmt.Printf("%d\t", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}
