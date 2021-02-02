package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Multiplication table for:", i)
		for j := 1; j <= 10; j++ {
			x := i * j
			fmt.Printf("%d * %d = %d\n", i, j, x)
		}
		fmt.Println("-----------------------------")
	}
}
