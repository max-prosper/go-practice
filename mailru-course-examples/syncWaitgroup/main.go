package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Println("goroutine ", in, "iteration ", j)
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}
