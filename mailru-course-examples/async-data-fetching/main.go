package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	result := make(chan string, 1)

	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")

		out <- "23 comments"
	}(result)

	return result
}

func getPage() {
	resultCh := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	commentsData := <-resultCh
	fmt.Println("main goroutine:", commentsData)

	return
}

func main() {
	getPage()
}
