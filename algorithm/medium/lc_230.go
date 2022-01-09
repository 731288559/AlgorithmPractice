package medium

// 230. 二叉搜索树中第K小的元素

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var counter int
	var result int
	var midOrder func(*TreeNode)

	midOrder = func(root *TreeNode) {
		if root.Left != nil {
			midOrder(root.Left)
		}
		counter++
		if counter == k {
			result = root.Val
			return
		}

		if root.Right != nil {
			midOrder(root.Right)
		}
	}
	midOrder(root)

	return result
}

func T_LC230() {
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

	println(kthSmallest(root, 3))

	root = &TreeNode{Val: 100}
	println(kthSmallest(root, 1))
}
