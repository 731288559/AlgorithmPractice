package leetcode

import (
	"fmt"
	"testing"
)

func TestLC3408_1(t *testing.T) {
	o := Constructor3408([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	o.Add(4, 104, 5)
	o.Edit(102, 8)
	fmt.Println(o.ExecTop())
	o.Rmv(101)
	o.Add(5, 105, 15)
	fmt.Println(o.ExecTop())
}

func TestLC3408_2(t *testing.T) {
	o := Constructor3408([][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}})
	o.Add(4, 104, 5)
	o.Edit(102, 9)
	fmt.Println(o.ExecTop())
	o.Rmv(101)
	o.Add(50, 101, 8)
	fmt.Println(o.ExecTop())
}

func TestLC3408_3(t *testing.T) {
	o := Constructor3408([][]int{{1, 26, 46}, {6, 27, 23}, {9, 24, 38}, {9, 20, 17}})
	o.Edit(24, 2)
	fmt.Println(o.ExecTop())
	fmt.Println(o.ExecTop())
	o.Rmv(20)
	o.Rmv(24)
	o.Add(0, 21, 31)
	o.Add(1, 4, 4)
	o.Add(9, 5, 4)
	o.Add(6, 17, 35)
	fmt.Println(o.ExecTop())
	o.Edit(4, 48)
	fmt.Println(o.ExecTop())
	fmt.Println(o.ExecTop())
}
