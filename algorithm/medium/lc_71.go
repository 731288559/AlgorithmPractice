package medium

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	for _, v := range strings.Split(path, "/") {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}

	}

	return "/" + strings.Join(stack, "/")
}

func T_LC71() {
	println(simplifyPath("/home/"), "/home")
	println(simplifyPath("/../"), "/")
	println(simplifyPath("/home//foo/"), "/home/foo")
	println(simplifyPath("/a/./b/../../c/"), "/c")
}
