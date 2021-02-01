package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func namedReturnPowerSeries(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a
	return
}

func main() {
	square, cube := powerSeries(2)
	fmt.Println("Square:", square, "Cube:", cube)

	square, cube = namedReturnPowerSeries(4)
	fmt.Println("Named Return:")
	fmt.Println("Square:", square, "Cube:", cube)
}
