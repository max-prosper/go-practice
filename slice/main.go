package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}

// func main() {
//
// 	xs := []int{1, 3, 5, 7, 9, 11}
//
// 	for i, value := range xs {
// 		fmt.Println(i, " - ", value)
// 	}
//
// }
//
//
// func main() {
//
// 	mySlice := make([]int, 0, 3)
//
// 	fmt.Println("-----------------")
// 	fmt.Println(mySlice)
// 	fmt.Println(len(mySlice))
// 	fmt.Println(cap(mySlice))
// 	fmt.Println("-----------------")
//
// 	for i := 0; i < 80; i++ {
// 		mySlice = append(mySlice, i)
// 		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
// 	}
// }
//
