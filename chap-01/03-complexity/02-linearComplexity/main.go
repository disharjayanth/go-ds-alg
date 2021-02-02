package main

import "fmt"

func main() {
	arr := [10]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 200
		fmt.Printf("Element [%d] = %d\n", i, arr[i])
	}
}
