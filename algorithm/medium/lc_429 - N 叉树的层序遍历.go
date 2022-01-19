package medium

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var q []*Node = []*Node{}
	for len(q) > 0 {
		n := len(q)
		var tmp = make([]int, n)
		for i := 0; i < n; i++ {
			cur := q[0]
			tmp[i] = cur.Val
			q = q[1:]
			for _, child := range cur.Children {
				q = append(q, child)
			}
		}
		res = append(res, tmp)
	}

	return res
}

func T_LC429() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}
	fmt.Println(levelOrder(root))
}
