package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker id: %d: Iteration: %d\n", id, i)
	}
}

func ExecuteWorker() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}