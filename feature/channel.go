package feature

import "sync"

// 两个goroutine交替打印01
func Channel_1() {
	n := 20
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= n; i = i + 2 {
			ch1 <- i
			if n, ok := <-ch2; ok {
				println(n)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= n; i = i + 2 {
			if n, ok := <-ch1; ok {
				println(n)
			}
			ch2 <- i
		}
		wg.Done()
	}()
	wg.Wait()
}
