package feature

import "fmt"

func SliceTest1() {
	m := make([]int, 5)
	m1 := m[1:4]
	fmt.Println(m1)
	m1[1] = 1
	m1 = append(m1, m...)
	fmt.Println(m1)
}

func SliceTest2() {
	var l []int
	l = append(l, 1)
	fmt.Println("cap:", cap(l))
	l = append(l, 2)
	fmt.Println("cap:", cap(l))
	l = append(l, 3)
	fmt.Println("cap:", cap(l))
}
