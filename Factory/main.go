package main

import (
	"fmt"
)

func main() {
	shape := GenerateShape("Square")
	if shape != nil {
		shape.Draw()
		fmt.Printf("Area: %f\n", shape.CalculateArea())
	}

	shape = GenerateShape("Circle")
	if shape != nil {
		shape.Draw()
		fmt.Printf("Area: %f\n", shape.CalculateArea())
	}
}
