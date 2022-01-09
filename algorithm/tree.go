package algorithm

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode) {
	fmt.Println(root.Val)
	if root.Left != nil {
		preOrder(root.Left)
	}
	if root.Right != nil {
		preOrder(root.Right)
	}
}

func midOrder(root *TreeNode) {
	if root.Left != nil {
		midOrder(root.Left)
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		midOrder(root.Right)
	}
}

func lastOrder(root *TreeNode) {
	if root.Left != nil {
		lastOrder(root.Left)
	}
	if root.Right != nil {
		lastOrder(root.Right)
	}
	fmt.Println(root.Val)
}

func T_order() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println("preOrder:")
	preOrder(root)
	fmt.Println("midOrder:")
	midOrder(root)
	fmt.Println("lastOrder:")
	lastOrder(root)
}
