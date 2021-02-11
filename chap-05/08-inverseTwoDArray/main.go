package main

import "fmt"

func determinant(m1 [2][2]int) (t float64) {
	det := 0
	det = det + ((m1[0][0] * m1[1][1]) - (m1[0][1] * m1[1][0]))
	return float64(det)
}

func inverse(m1 [2][2]int) (r [2][2]int) {
	det := determinant(m1)
	invMatrix := [2][2]int{}
	invMatrix[0][0] = m1[1][1] / int(det)
	invMatrix[0][1] = -1 * m1[0][1] / int(det)
	invMatrix[1][0] = -1 * m1[0][1] / int(det)
	invMatrix[1][1] = m1[0][0] / int(det)
	return invMatrix
}

func main() {
	m1 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	inv := inverse(m1)
	for _, row := range inv {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}
