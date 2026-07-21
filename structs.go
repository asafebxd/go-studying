package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

type Employee struct {
	Name   string
	Salary float64
	Active bool
}

func (product Product) IsAvailable() bool {
	//	return product.Stock > 0 also works

	if product.Stock <= 0 {
		return false
	}

	return true
}

func (product *Product) ApplyDiscount(percent float64) {
	product.Price = product.Price - (product.Price * percent / 100)
}

func NewEmployee(name string, salary float64) Employee {
	return Employee{
		Name:   name,
		Salary: salary,
		Active: true,
	}
}

func (employee *Employee) Deactivate() {
	employee.Active = false
}

func strcuts() {
	product := Product{
		Name:  "Keyboard",
		Price: 200,
		Stock: 10,
	}

	product.ApplyDiscount(10)

	fmt.Println(product.Price)
}
