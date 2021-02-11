package main

import "fmt"

func transpose(m1 [4][4]int) (t [4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			t[i][j] = m1[j][i]
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

	t := transpose(m1)
	for _, row := range t {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
