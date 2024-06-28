package main

import "fmt"

type IShape interface {
	Draw()
	CalculateArea() float64
}

type Square struct {
	Side float64
}

func (sq *Square) Draw() {
	fmt.Println("Drawing square")
}

func (sq *Square) CalculateArea() float64 {
	return sq.Side * sq.Side
}

type Circle struct {
	Radius float64
}

func (c *Circle) Draw() {
	fmt.Println("Drawing circle")
}

func (c *Circle) CalculateArea() float64 {
	return 3.14 * c.Radius * c.Radius
}
