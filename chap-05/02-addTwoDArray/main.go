package main

import "fmt"

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	sum := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	return sum
}

func main() {
	m1 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	m2 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	sum := add(m1, m2)
	for _, row := range sum {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
