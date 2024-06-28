package main

func GenerateShape(shapeType string) IShape {
	if shapeType == "Square" {
		return &Square{Side: 4}
	}
	if shapeType == "Circle" {
		return &Circle{Radius: 3}
	}
	return nil
}
