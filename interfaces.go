package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Height float64
}

const pi = 3.14159

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Area() float64 {
	return pi * math.Pow(circle.Radius, 2)
}

func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

type PaymentProcessor interface {
	Process(amount float64) string	
}

type CreditCard struct {}
type Pix struct{}

func (CreditCard) Process(amount float64) string {
	return fmt.Sprintf("Processed %0.2f using credit card", amount)
}

func (Pix) Process(amount float64) string {
	return fmt.Sprintf("Processed %0.2f using Pix", amount)
}

func completePayment(processor PaymentProcessor, amount float64) string {
	return processor.Process(amount)
}

func identifyType(value any) string {
	switch value.(type) {
	case string:
		return "string"
	case int:
		return "integer"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}


func interfaces() {

}