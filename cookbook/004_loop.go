package main

import "fmt"

func main() {
	fmt.Printf("004_loop")

	// number loop
	fmt.Printf("number loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// list object loop
	fmt.Printf("list object loop")
	nums := []int{100, 99, 98}
	for index, item := range nums {
		fmt.Println(index, item)
	}

}
