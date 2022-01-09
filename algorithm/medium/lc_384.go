package medium

import (
	"fmt"
	"math/rand"
)

// 384. 打乱数组

type Solution struct {
	nums []int
}

func Constructor384(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	l := make([]int, n)
	copy(l, this.nums)
	for i := range l {
		j := i + rand.Intn(n-i)
		l[i], l[j] = l[j], l[i]
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func T_LC384() {
	obj := Constructor384([]int{1, 2, 3})
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Reset())
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Shuffle())
}
