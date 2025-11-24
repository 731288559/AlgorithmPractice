package leetcode

import (
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	nodeList := []*Node{
		{Val: 1},
		{Val: 2},
		{Val: 3},
	}
	nodeList[0].Next = nodeList[1]
	nodeList[1].Next = nodeList[2]
	nodeList[0].Random = nodeList[2]
	printList(nodeList[0])

	printList(copyRandomList(nodeList[0]))
}
