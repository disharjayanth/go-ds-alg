package main

import "fmt"

func multiply(matrix1 [4][4]int, matrix2 [4][4]int) (mul [4][4]int) {
	for i, row := range matrix1 {
		for j, col := range row {
			mul[i][j] = col * matrix2[i][j]
		}
	}
	return mul
}

func main() {
	m1 := [4][4]int{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}

	m2 := [4][4]int{
		{3, 3, 3, 3},
		{3, 6, 3, 3},
		{3, 3, 3, 3},
		{3, 3, 3, 3},
	}

	mul := multiply(m1, m2)

	for _, row := range mul {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
