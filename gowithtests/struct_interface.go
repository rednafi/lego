package main

import "math"

//import "fmt"

// Rectangle defines a rectangle type consisting of height and width
type Rectangle struct {
	length float64
	width  float64
}

// Circle defines a circle type that has radius
type Circle struct {
	radius float64
}

// Shape is an interface that binds Area() with Rectangle & Circle
type Shape interface {
	Area() float64
}

// Perimeter function calculates the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2.0 * (rectangle.length + rectangle.width)
}

// Area method calculates the area of a rectangle
func (rectangle Rectangle) Area() float64 {
	return (rectangle.length * rectangle.width)
}

// Area method also calculates the area of a rectangle
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius
}

// func main() {
// 	fmt.Println("Perimeter:", Perimeter(Rectangle{10, 10}))

// 	type Shape struct {
// 		circle    Circle
// 		rectangle Rectangle
// 	}
// 	shape := Shape{Circle{10}, Rectangle{10, 10}}

// 	fmt.Println("Area:", shape.circle.Area())
// 	fmt.Println("Area:", shape.rectangle.Area())
// }
