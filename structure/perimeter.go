package structure

import "math"

/**
Suppose that we need some geometry code to calculate the perimeter of a rectangle given a height and width.
We can write a Perimeter(width float64, height float64) function, where float64 is for floating-point numbers like 123.45.
**/

type Rectangle struct {
	width  float64
	height float64
}

func (p Rectangle) Area() float64 {
	return (p.width * p.height) 
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64{
	return (t.Base*t.Height)*0.5
}
type Shape interface{
	Area() float64
}