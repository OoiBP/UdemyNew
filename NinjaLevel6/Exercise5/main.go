package main

import (
	"fmt"
	"math"
)

type square struct {
	L float64
	W float64
}
type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.L * s.W
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		L: 3.1,
		W: 4.3,
	}
	c1 := circle{
		r: 20,
	}
	info(s1)
	info(c1)
}

/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/
