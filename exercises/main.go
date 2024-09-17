package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

// you can say that with these methods we kind of can keep composing the struct with new methods as required, these methods are bound to the type
func (c Circle) calculateCircumference() float64 {
	return 2 * PI * c.radius
}

func (c Circle) calculateArea() float64 {
	return PI * c.radius * c.radius
}

func main() {
	myCircle := NewCircle(2)

	fmt.Printf("The circumference of circle with radius %f is %f and area is %f\n", myCircle.radius, myCircle.calculateCircumference(), myCircle.calculateArea())
}
