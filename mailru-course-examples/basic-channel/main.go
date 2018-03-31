package main

import (
	"fmt"
)

func main() {
	in := make(chan int, 0)

	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			fmt.Println("before", i)
			out <- 1
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}

	fmt.Scanln()
}

// func main() {
// 	ch1 := make(chan int)

// 	go func(in chan int) {
// 		val := <-in
// 		fmt.Println("GO: get from chan", val)
// 		fmt.Println("GO: after read from chan")
// 	}(ch1)

// 	ch1 <- 42

// 	fmt.Println("MAIN: after put to chan")
// 	fmt.Scanln()
// }
