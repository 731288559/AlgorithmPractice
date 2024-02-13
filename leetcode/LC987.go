package leetcode

import (
	"math"
	"sort"
)

/*
987. 二叉树的垂序遍历

给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。

对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。

二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。

返回二叉树的 垂序遍历 序列。
*/
func verticalTraversal(root *TreeNode) [][]int {
	type P struct {
		Row int
		Col int
		Val int
	}
	var l []P

	var dfs func(n *TreeNode, row, col int)
	dfs = func(n *TreeNode, row, col int) {
		if n == nil {
			return
		}
		l = append(l, P{Row: row, Col: col, Val: n.Val})
		dfs(n.Left, row+1, col-1)
		dfs(n.Right, row+1, col+1)
	}

	dfs(root, 0, 0)

	sort.Slice(l, func(i, j int) bool {
		a, b := l[i], l[j]
		if a.Col != b.Col {
			return a.Col < b.Col
		}
		if a.Row != b.Row {
			return a.Row < b.Row
		}
		return a.Val < b.Val
	})

	var ans [][]int

	lastCol := math.MinInt
	for _, n := range l {
		if n.Col != lastCol {
			ans = append(ans, nil)
			lastCol = n.Col
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], n.Val)
	}

	return ans
}
