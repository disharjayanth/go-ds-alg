package main

import "fmt"

func substract(matrix1 [4][4]int, matrix2 [4][4]int) (subs [4][4]int) {
	subs = [4][4]int{}
	for i, row := range matrix2 {
		for j, col := range row {
			subs[i][j] = matrix1[i][j] - col
		}
	}
	return
}

func main() {
	m1 := [4][4]int{
		{5, 5, 5, 5},
		{5, 5, 5, 5},
		{5, 5, 5, 5},
		{5, 5, 5, 5},
	}

	m2 := [4][4]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	subs := substract(m1, m2)
	for _, row := range subs {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
