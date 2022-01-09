package easy

// 559. N 叉树的最大深度

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	// var dfs func(n *Node) int
	// dfs = func(node *Node) int {
	// 	if node == nil {
	// 		return 0
	// 	}
	// 	if len(node.Children) == 0 {
	// 		return 1
	// 	}
	// 	max := 0
	// 	tmp := 0
	// 	for _, n := range node.Children {
	// 		tmp = dfs(n) + 1
	// 		if tmp > max {
	// 			max = tmp
	// 		}
	// 	}
	// 	return max
	// }
	// return dfs(root)

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
			&Node{Val: 3, Children: []*Node{
				&Node{Val: 5},
				&Node{Val: 6},
			}},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}
	println(maxDepth(root), 3)
}
