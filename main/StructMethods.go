package main

import "fmt"


type rectangle struct {
	width, height int
}

func (r *rectangle) makeSquarePtr() {
	r.width = r.height
}


func (r rectangle) makeSquare() {
	r.width = r.height
}

func (r rectangle) printProperties() {
	fmt.Println("Width: ", r.width)
	fmt.Println("Height: ", r.height)
}

func main() {
	r := rectangle{width: 10, height: 20}

	fmt.Println("Initial rectangle: ")
	r.printProperties()

	r.makeSquare()
	fmt.Println("After makeSquare: ")
	r.printProperties()

	r.makeSquarePtr()
	fmt.Println("After makeSquarePtr: ")
	r.printProperties()
}
