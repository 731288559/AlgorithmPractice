package leetcode

/*
	2641. 二叉树的堂兄弟节点 II

	给你一棵二叉树的根 root ，请你将每个节点的值替换成该节点的所有 堂兄弟节点值的和 。

	如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 堂兄弟 。

	请你返回修改值之后，树的根 root 。

	注意，一个节点的深度指的是从树根节点到这个节点经过的边数。
*/

func replaceValueInTree(root *TreeNode) *TreeNode {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		sum := 0
		q2 := make([]*TreeNode, 0)

		for _, n := range q {
			if n.Left != nil {
				q2 = append(q2, n.Left)
				sum += n.Left.Val
			}
			if n.Right != nil {
				q2 = append(q2, n.Right)
				sum += n.Right.Val
			}
		}

		for _, n := range q {
			childSum := 0
			if n.Left != nil {
				childSum += n.Left.Val
			}
			if n.Right != nil {
				childSum += n.Right.Val
			}
			if n.Left != nil {
				n.Left.Val = sum - childSum
			}
			if n.Right != nil {
				n.Right.Val = sum - childSum
			}
		}

		q = q2
	}
	root.Val = 0
	return root
}
