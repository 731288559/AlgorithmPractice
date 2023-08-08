package leetcode

import "math"

/*
	822. 翻转卡片游戏

	在桌子上有 N 张卡片，每张卡片的正面和背面都写着一个正数（正面与背面上的数有可能不一样）。
	我们可以先翻转任意张卡片，然后选择其中一张卡片。
	如果选中的那张卡片背面的数字 X 与任意一张卡片的正面的数字都不同，那么这个数字是我们想要的数字。
	哪个数是这些想要的数字中最小的数（找到这些数中的最小值）呢？如果没有一个数字符合要求的，输出 0。
	其中, fronts[i] 和 backs[i] 分别代表第 i 张卡片的正面和背面的数字。
	如果我们通过翻转卡片来交换正面与背面上的数，那么当初在正面的数就变成背面的数，背面的数就变成正面的数。

	示例：

	输入：fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
	输出：2
	解释：假设我们翻转第二张卡片，那么在正面的数变成了 [1,3,4,4,7] ， 背面的数变成了 [1,2,4,1,3]。
	接着我们选择第二张卡片，因为现在该卡片的背面的数是 2，2 与任意卡片上正面的数都不同，所以 2 就是我们想要的数字。
*/

func flipgame(fronts []int, backs []int) int {
	// 1 2 4 4 7
	// 1 3 4 1 3

	/*
		1 2 4 4 7
		1 3 4 1 3

		1 2 4 1 3
		1 3 4 4 7
	*/
	ans := math.MaxInt32

	m := make(map[int]struct{}, len(fronts))
	for i, n := range fronts {
		if n == backs[i] {
			m[n] = struct{}{}
		}
	}

	for i, n := range fronts {
		if n < ans {
			if _, ok := m[n]; !ok {
				ans = n
			}
		}

		back := backs[i]
		if back < ans {
			if _, ok := m[back]; !ok {
				ans = back
			}
		}
	}
	if ans == math.MaxInt32 {
		ans = 0
	}
	return ans
}
