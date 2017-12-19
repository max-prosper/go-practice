package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello wodrld!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet)
}
