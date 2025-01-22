package concurrency

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func ExecuteSampleGoSched() {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("trace2.out")
	trace.Start(f)
	defer trace.Stop()
	var wg sync.WaitGroup
	wg.Add(2)
	go funcA(&wg)
	go funcB(&wg)
	wg.Wait()
}

func funcA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("A%d\n", i)
		//runtime.Gosched()
	}
}
func funcB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("B%d\n",i)
		//runtime.Gosched()
	}
}