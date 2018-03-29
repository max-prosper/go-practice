package main

import (
	"fmt"
	"runtime"
)

const (
	iterationNum  = 7
	goroutinesNum = 5
)

func doSomeWork(in int) {
	for j := 0; j < iterationNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in, j int) string {
	return fmt.Sprintln("routine", in, "iteration", j)
}

func main() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}

	fmt.Scanln()
}
