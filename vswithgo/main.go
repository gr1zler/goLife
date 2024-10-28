package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float32
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * float64(s.width)
}
func (c *circle) change() {
	c.radius = 2
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	c := circle{3}
	c.change()
	fmt.Print(c.radius)
}
