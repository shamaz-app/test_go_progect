package main

import "fmt"
import "math"

type figure interface {
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) perim() float64 {
	return r.width * r.height
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(f figure) {
	fmt.Println(f)
	fmt.Println(f.perim())
}

func main() {
	r := rect{width: 10, height: 20}
	c := circle{radius: 3}

	measure(r)
	measure(c)
}