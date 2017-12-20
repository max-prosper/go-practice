package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0xc42007c008

	changeMe(&age)

	fmt.Println(&age) // 0xc42007c008
	fmt.Println(age)  // 24
}

func changeMe(z *int) {
	fmt.Println(z)  // 0xc42007c008
	fmt.Println(*z) // 24
	*z = 24
	fmt.Println(z)  // 0xc42007c008
	fmt.Println(*z) // 24
}
