package easy

// 1436. 旅行终点站

func destCity(paths [][]string) string {
	m := make(map[string]string)
	for i := range paths {
		m[paths[i][0]] = paths[i][1]
	}

	dest := paths[0][1]
	for {
		if _, ok := m[dest]; !ok {
			return dest
		}
		dest = m[dest]
	}
}

func destCity2(paths [][]string) string {
	m := make(map[string]struct{})
	for i := range paths {
		m[paths[i][1]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := m[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
