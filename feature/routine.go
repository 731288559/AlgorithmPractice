package feature

import (
	"sync"
	"time"
)

func Routine_1() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		wg.Done()
	}
	wg.Wait()
	println("main exit")
}
