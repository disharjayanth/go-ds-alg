package main

import "fmt"

func printZigZag(m1 [4][4]int) {
	m := 4
	n := 4
	for k := 0; k < m-1; k++ {
		i := k
		j := 0
		for i >= 0 {
			fmt.Printf("%d\t", m1[i][j])
			i = i - 1
			j = j + 1
		}
		fmt.Println()
	}

	for k := 1; k <= n-1; k++ {
		i := m - 1
		j := k
		for j <= n-1 {
			fmt.Printf("%d\t", m1[i][j])
			i = i - 1
			j = j + 1
		}
		fmt.Println()
	}
}

func main() {
	m1 := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	printZigZag(m1)
}
