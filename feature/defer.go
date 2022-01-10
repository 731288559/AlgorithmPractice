package feature

import "fmt"

func f1() bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	var m map[int]bool
	m[1] = true

	return true
}

func DeferTest() {
	fmt.Println("f1:", f1())
}
