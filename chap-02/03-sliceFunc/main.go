package main

import "fmt"

func twiceValue(arr []int) {
	for i, num := range arr {
		arr[i] = 2 * num
	}
}

func display(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	arr := []int{1, 3, 5, 6}
	twiceValue(arr)

	for _, n := range arr {
		fmt.Println(n)
	}

	display()
}
