package leetcode

/*
103. 二叉树的锯齿形层序遍历

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	fromLeft := true
	for len(q) > 0 {
		var q2 []*TreeNode
		var tmp []int
		for _, n := range q {
			if n.Left != nil {
				q2 = append(q2, n.Left)
			}
			if n.Right != nil {
				q2 = append(q2, n.Right)
			}
			tmp = append(tmp, n.Val)
		}
		if !fromLeft {
			for i, n := 0, len(tmp); i < n/2; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
		}

		fromLeft = !fromLeft
		q = q2
		res = append(res, tmp)
	}
	return res
}
