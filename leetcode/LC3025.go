package leetcode

/*
	3025. 人员站位的方案数 I

	给你一个  n x 2 的二维数组 points ，它表示二维平面上的一些点坐标，其中 points[i] = [xi, yi] 。

	计算点对 (A, B) 的数量，其中A 在 B 的左上角，并且它们形成的长方形中（或直线上）没有其它点（包括边界）。
	返回数量。
*/

// 暴力解法
func numberOfPairs(points [][]int) int {
	ans := 0

	n := len(points)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			if !(x1 <= x2 && y1 >= y2) {
				continue
			}

			ok := true
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				x3, y3 := points[k][0], points[k][1]
				if x1 <= x3 && x3 <= x2 && y2 <= y3 && y3 <= y1 {
					ok = false
				}
			}
			if ok {
				ans++
			}
		}
	}

	return ans
}

// 优化：先排序后遍历，记录处理情况剪枝
