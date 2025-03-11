package instruct

import (
	"fmt"
	"math"
)

// Perimeter takes the width and height of a rectangle and returns the perimeter.
func Perimeter(width float64, height float64) (perimeter float64) {
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)
	return
}

// Area takes the width and height of a rectangle and returns the area.
func Area(width float64, height float64) (area float64) {
	area = width * height
	fmt.Println(area)
	return
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rec Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (rec.Width + rec.Height)
	fmt.Println(perimeter)
	return
}

func (rec Rectangle) Area() (area float64) {
	area = rec.Width * rec.Height
	fmt.Println(area)
	return
}

func (triangle Triangle) Area() (area float64) {
	area = 0.5 * (triangle.Base * triangle.Height)
	fmt.Println(area)
	return
}

func (circle Circle) Area() (area float64) {
	area = circle.Radius * circle.Radius * math.Pi
	fmt.Println(area)
	return
}

type Shape interface {
	Area() float64
}
