package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct{}

func (ds DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shaping.... drawShap method from DrawShape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (dc DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	dc.shape.drawShape(dc.x, dc.y)
}

func (dc DrawContour) resizeByFactor(factor int) {
	dc.factor = factor
}

func main() {
	x := [5]float32{1, 2, 3, 4, 5}
	y := [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
