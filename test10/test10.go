package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Cricle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Cricle) area() float64 {
	return math.Pi*c.radius + c.radius
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	cricle := Cricle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{h: 3, w: 4}
	fmt.Println(getArea(cricle))
	fmt.Println(getArea(rectangle))
}
