package main

import "fmt"

func main() {
	twoDArray := [8][8]int{}
	twoDArray[3][6] = 18
	twoDArray[7][4] = 3
	for i, arr := range twoDArray {
		fmt.Println("Array no:", i, arr)
	}

	// 2d arrays rows and cols
	rows := 7
	cols := 9
	twoDSlices := make([][]int, rows)

	for i := range twoDSlices {
		twoDSlices[i] = make([]int, cols)
	}

	fmt.Println("twoDSlices:", twoDSlices)

	arr := []int{5, 6, 7, 8, 9}
	slice1 := arr[:3]
	fmt.Println("Slice1:", slice1)

	slice2 := arr[1:5]
	fmt.Println("Slice2:", slice2)

	slice3 := append(slice2, 12)
	fmt.Println("Slice3:", slice3)

	slice4 := append(arr, slice3...)
	fmt.Println("Slice4:", slice4)
}
