package main

import "fmt"

func main() {
	m := [10]int{}

	for i := 0; i < len(m); i++ {
		m[i] = i + 200
		fmt.Printf("Element [%d] = %d\n", i, m[i])
	}
}
