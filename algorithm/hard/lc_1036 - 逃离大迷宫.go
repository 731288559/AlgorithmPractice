package hard

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	maxSize := len(blocked) * (len(blocked) - 1) / 2
	var boundary int = 1e6

	type Point struct {
		X int
		Y int
	}

	blockMap := make(map[Point]struct{})
	for _, block := range blocked {
		blockMap[Point{X: block[0], Y: block[1]}] = struct{}{}
	}

	check := func(a, b Point) bool {
		l := []Point{a}
		visited := make(map[Point]struct{})
		visited[a] = struct{}{}
		for len(l) > 0 && len(visited) <= maxSize {
			curP := l[0]
			l = l[1:]
			for _, dir := range dirs {
				p := Point{X: curP.X + dir[0], Y: curP.Y + dir[1]}
				if p.X < 0 || p.X >= boundary || p.Y < 0 || p.Y >= boundary {
					continue
				}
				if _, ok := visited[p]; ok {
					continue
				}
				if _, ok := blockMap[p]; ok {
					continue
				}
				if p.X == b.X && p.Y == b.Y {
					return true
				}
				visited[p] = struct{}{}
				l = append(l, p)
			}
		}
		return len(visited) > maxSize
	}
	a := Point{source[0], source[1]}
	b := Point{target[0], target[1]}
	return check(a, b) && check(b, a)
}

func isEscapePossible2(block [][]int, source, target []int) bool {
	type pair struct{ x, y int }
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	const (
		blocked = -1 // 在包围圈中
		valid   = 0  // 不在包围圈中
		found   = 1  // 无论在不在包围圈中，但在 n(n-1)/2 步搜索的过程中经过了 target

		boundary int = 1e6
	)

	n := len(block)
	if n < 2 {
		return true
	}

	blockSet := map[pair]bool{}
	for _, b := range block {
		blockSet[pair{b[0], b[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		countdown := n * (n - 1) / 2

		q := []pair{{sx, sy}}
		vis := map[pair]bool{{sx, sy}: true}
		for len(q) > 0 && countdown > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				np := pair{x, y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
					if x == fx && y == fy {
						return found
					}
					countdown--
					vis[np] = true
					q = append(q, np)
				}
			}
		}

		if countdown > 0 {
			return blocked
		}
		return valid
	}

	res := check(source, target)
	return res == found || res == valid && check(target, source) != blocked
}

func T_LC1036() {
	println(isEscapePossible([][]int{{0, 1}, {1, 0}}, []int{0, 0}, []int{0, 2}), false)
	println(isEscapePossible([][]int{}, []int{0, 0}, []int{999999, 999999}), true)
}
