package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area()
	perimeter()
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() {
	fmt.Println(r.width * r.height)
}

func (r rect) perimeter() {
	fmt.Println(2*r.width + 2*r.height)
}

func (c circle) perimeter() {
	fmt.Println(2 * math.Pi * c.radius)
}

func (c circle) area() {
	fmt.Println(math.Pi * c.radius * c.radius)
}

func measurement(g geometry) {
	g.area()
	g.perimeter()
}
func main() {

	r := rect{width: 9.1, height: 5.5}
	c := circle{radius: 10}

	measurement(r)
	measurement(c)
}
