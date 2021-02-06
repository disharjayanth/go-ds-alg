package main

import "fmt"

func h(x int, y int) int {
	return x * y
}

func g(l int, m int) (x int, y int) {
	x = l * 2
	y = m * 4
	return
}

func main() {
	fmt.Println(h(2, 2))
	fmt.Println(g(2, 2))
}
