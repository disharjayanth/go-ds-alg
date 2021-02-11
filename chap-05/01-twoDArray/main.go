package main

import (
	"fmt"
)

var arr = [4][5]int{
	{4, 5, 7, 8, 9},
	{1, 2, 4, 5, 6},
	{9, 10, 11, 12, 14},
	{3, 5, 6, 8, 9},
}

func main() {
	for _, row := range arr {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}

	fmt.Println("****************************************************")
	fmt.Println("Element at row 2 and col 3:", arr[2][3])
	fmt.Println("****************************************************")

	fmt.Println("Row matrix:")
	matrix := [1][3]int{
		{1, 2, 3},
	}
	fmt.Println(matrix)

	fmt.Println("****************************************************")
	fmt.Println("Column matrix:")
	matrix2 := [4][1]int{
		{1},
		{2},
		{3},
		{4},
	}
	fmt.Println(matrix2)

	fmt.Println("****************************************************")
	fmt.Println("Lower traingular matrix:")
	matrix3 := [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}
	for _, row := range matrix3 {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}

	fmt.Println("****************************************************")
	fmt.Println("Upper traingular matrix:")
	matrix4 := [3][3]int{
		{1, 2, 3},
		{0, 1, 4},
		{0, 0, 5},
	}
	for _, row := range matrix4 {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}

	fmt.Println("****************************************************")
	fmt.Println("Null matrix:")
	matrix5 := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	for _, row := range matrix5 {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}

	fmt.Println("****************************************************")
	fmt.Println("Identity matrix:")

	matrix6 := identityMatrix(4)
	for _, row := range matrix6 {
		for _, col := range row {
			fmt.Printf("%f\t", col)
		}
		fmt.Println()
	}
}

func identityMatrix(order int) [][]float64 {
	matrix := make([][]float64, order)
	for i, _ := range matrix {
		temp := make([]float64, order)
		matrix[i] = temp
		for j, _ := range temp {
			if i == j {
				matrix[i][j] = 1.0
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
