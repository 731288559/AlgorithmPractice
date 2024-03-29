package leetcode

// 337. 打家劫舍 III

//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
//除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
//给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

//输入: root = [3,2,3,null,3,null,1]
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
//输入: root = [3,4,5,1,3,null,1]
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9

func rob3(root *TreeNode) int {
	var dfs func(n *TreeNode) (int, int) // 返回取和不取当前节点时的最大价值
	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		leftGetRoot, leftNoRoot := dfs(n.Left)
		rightGetRoot, rightNoRoot := dfs(n.Right)
		return n.Val + leftNoRoot + rightNoRoot, max(leftGetRoot, leftNoRoot) + max(rightGetRoot, rightNoRoot)
	}

	return max(dfs(root))
}
