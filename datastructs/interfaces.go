package main

import "fmt"

// AreaPerim the main interface
type AreaPerim interface{
	Area() float64
	Perim() float64
}

// Circle is a struct that AreaPerim works on
type Circle struct {
	name string
	radius float64
}

// Rectangle is a struct that AreaPerim works on
type Rectangle struct {
	name string
	length float64
	width float64
}

// Area defined for Circle
func (c *Circle) Area() float64 {
	return c.radius*c.radius
}

// Perim defined for Circle
func (c *Circle) Perim() float64 {
	return 2*3.1416*c.radius
}

// Area defined for Rectangle
func (r *Rectangle) Area() float64{
	return r.length*r.width
}

// Perim defined for Rectangle
func (r *Rectangle) Perim() float64{
	return 2*(r.length + r.width)
}


func main() {
	c := Circle{name: "Circle", radius:10.0}
	r := Rectangle{name: "Rectangle", length:10, width:20}

	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perim:", c.Perim())
	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perim:", r.Perim())

}
