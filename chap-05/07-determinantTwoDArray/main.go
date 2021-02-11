package main

import "fmt"

func determinant(m1 [2][2]int) float32 {
	det := 0
	det = det + ((m1[0][0] * m1[1][1]) - (m1[0][1] * m1[1][0]))
	return float32(det)
}

func main() {
	m1 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	det := determinant(m1)
	fmt.Println(det)
}
