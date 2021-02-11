package main

import "fmt"

func multiply(m1 [4][4]int, m2 [4][4]int) (mul [4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mulSum := 0
			for n := 0; n < 4; n++ {
				mulSum = mulSum + m1[i][n]*m2[n][j]
			}
			mul[i][j] = mulSum
		}
	}
	return
}

func main() {
	m1 := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	m2 := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	mul := multiply(m1, m2)
	for _, row := range mul {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
