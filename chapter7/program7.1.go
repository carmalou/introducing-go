package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x1: 0, y1: 0, x2: 0, y2: 0, length: 3, width: 4}

	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1     float64
	y1     float64
	x2     float64
	y2     float64
	length float64
	width  float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length * r.width)
}

type Shape interface {
	area() float64
	perimeter() float64
}
