package easy

// 563. 二叉树的坡度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	sum := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sumLeft := f(root.Left)
		sumRight := f(root.Right)
		n := sumLeft - sumRight
		if n < 0 {
			n = -n
		}
		sum += n
		return root.Val + sumLeft + sumRight
	}
	f(root)
	return sum
}

func T_LC563() {
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	println(findTilt(&root), 1)
	root = TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7}},
	}
	println(findTilt(&root), 15)
}
