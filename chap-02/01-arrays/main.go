package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("for-loop")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, ":", arr[i])
	}

	fmt.Println("for-range")
	for i, n := range arr {
		fmt.Println(i, ":", n)
	}
}
