package main

import "fmt"

type areaPerim interface{
	area() float64
	perim() float64
}

type circle struct {
	name string
	radius float64
}

type rectangle struct {
	name string
	length float64
	width float64
}

func (c *circle) area() float64 {
	return c.radius*c.radius
}

func (c *circle) perim() float64 {
	return 2*3.1416*c.radius
}

func (r *rectangle) area() float64{
	return r.length*r.width
}

func (r *rectangle) perim() float64{
	return 2*(r.length + r.width)
}


func main() {
	c := circle{name: "circle", radius:10.0}
	r := rectangle{name: "rectangle", length:10, width:20}

	fmt.Println("circle area:", c.area())
	fmt.Println("circle perim:", c.perim())
	fmt.Println("rectangle area:", r.perim())
	fmt.Println("rectangle perim:", r.perim())

}
