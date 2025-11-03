package feature

import "golang.org/x/sync/semaphore"

func SemaphoreTest() {
	semaphore.NewWeighted(1)
}
