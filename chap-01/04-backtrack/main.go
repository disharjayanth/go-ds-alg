package main

import "fmt"

func findElementsWithSum(arr [10]int, combinations [6]int, size int, addedSum int, addValue int, l int, m int) int {
	num := 0

	if addValue > addedSum {
		return -1
	}

	if addValue == addedSum {
		num = num + 1
		for i := 0; i < m; i++ {
			fmt.Printf("%d,", arr[combinations[i]])
		}
		fmt.Println(" ")
	}

	for j := l; j < size; j++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, addedSum, addValue+arr[j], l, m+1)
		l = l + 1
	}

	fmt.Println("addValue:", addValue, "l:", l, "m:", m, "num:", num)
	return num
}

func main() {
	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	addedSum := 5
	combinations := [6]int{}

	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}
