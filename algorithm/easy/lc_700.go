package easy

// 700. 二叉搜索树中的搜索

func searchBST(root *TreeNode, val int) *TreeNode {
	n := root
	for n != nil {
		if n.Val == val {
			return n
		}
		if n.Val > val {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return nil
}
