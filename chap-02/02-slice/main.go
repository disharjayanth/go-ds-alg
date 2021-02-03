package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	for i, n := range arr {
		fmt.Println("index:", i, n)
	}
	fmt.Println("Length of arr:", len(arr))
	fmt.Println("Capacity of arr:", cap(arr))

	arr = append(arr, 5, 6, 7, 8, 10)
	for i, n := range arr {
		fmt.Println("index:", i, n)
	}
	fmt.Println("Length of arr:", len(arr))
	fmt.Println("Capacity of arr:", cap(arr))
}
