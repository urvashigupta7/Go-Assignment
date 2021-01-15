package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) Area() float64 {
	return rect.length * rect.breadth
}

func (circ Circle) Area() float64 {
	return 3.14 * circ.radius * circ.radius
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.length + rect.breadth)
}

func (circ Circle) Perimeter() float64 {
	return 2 * 3.14 * (circ.radius)
}

func main() {
	shape := os.Args[1]

	if shape == "circle" {
		radius := os.Args[2]
		r, _ := strconv.ParseFloat(radius, 64)
		circle := Circle{r}
		area := circle.Area()
		perimeter := circle.Perimeter()
		fmt.Println(area)
		fmt.Println(perimeter)
	} else {
		l := os.Args[2]
		b := os.Args[3]
		length, _ := strconv.ParseFloat(l, 64)
		breadth, _ := strconv.ParseFloat(b, 64)
		rectangle := Rectangle{length, breadth}
		area := rectangle.Area()
		perimeter := rectangle.Perimeter()
		fmt.Println(area)
		fmt.Println(perimeter)
	}

}
