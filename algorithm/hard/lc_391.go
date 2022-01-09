package hard

import (
	"sort"
)

// 391. 完美矩形

type LineSlice [][]int

func (l LineSlice) Len() int      { return len(l) }
func (l LineSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l LineSlice) Less(i, j int) bool {
	if l[i][0] != l[j][0] {
		return l[i][0]-l[j][0] < 0
	}
	return l[i][1] < l[j][1]
}

func isRectangleCover(rectangles [][]int) bool {
	lines := LineSlice(make([][]int, len(rectangles)*2))
	for i, n := range rectangles {
		// n = {x, y, a, b}
		lines[2*i] = []int{n[0], n[3], n[1], -1}  // 左竖线
		lines[2*i+1] = []int{n[2], n[3], n[1], 1} // 右竖线
	}
	// 根据横坐标和纵坐标排序，竖线横坐标和竖线上顶点小的排前面
	sort.Sort(lines)

	var (
		l1 [][]int
		l2 [][]int
	)
	for i := 0; i < len(lines); {
		l1, l2 = [][]int{}, [][]int{}
		// 合并x相同的线
		right := len(lines)
		for j := i; j < len(lines); j++ {
			cur := lines[j]
			if cur[0] != lines[i][0] {
				right = j
				break
			}
			var l [][]int
			// TODO: 这个地方实现得比较丑陋，想办法改进一下
			if cur[3] == -1 {
				l = l1
			} else {
				l = l2
			}

			if len(l) == 0 {
				l = append(l, cur)
			} else {
				prev := l[len(l)-1]
				if prev[0] != cur[0] { // x不同
					l = append(l, cur)
					continue
				}
				if prev[1] == cur[2] { // 恰好相接
					l[len(l)-1] = []int{prev[0], cur[1], prev[2], prev[3]}
				} else if prev[1] < cur[2] { // 相隔
					l = append(l, cur)
				} else { // 相交
					return false
				}
			}

			if cur[3] == -1 {
				l1 = l
			} else {
				l2 = l
			}
		}
		if i > 0 && right < len(lines) { // 内部竖线
			if len(l1) != len(l2) {
				return false
			}
			for k := range l1 {
				if l1[k][1] == l2[k][1] && l1[k][2] == l2[k][2] {
					continue
				}
				return false
			}
		} else { // 两边竖线
			if len(l1)+len(l2) != 1 {
				return false
			}
		}
		i = right
	}

	return true
}

func T_LC391() {
	println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4},
	}), true)

	println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4},
	}), false)

	println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4},
	}), false)
}
