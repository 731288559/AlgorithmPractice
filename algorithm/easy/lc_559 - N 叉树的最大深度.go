package easy

// 559. N 叉树的最大深度

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		curChildDepth := maxDepth(child)
		if curChildDepth > maxChildDepth {
			maxChildDepth = curChildDepth
		}
	}
	return maxChildDepth + 1
}

func T_LC559() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	println(maxDepth(root), 3)
}
