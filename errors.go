package main

import "errors"

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func validateProduct(name string, price float64, stock int) error {
	if name == "" {
		return errors.New("name is required")
	}

	if price <= 0 {
		return errors.New("price must be greater than zero") 
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return nil
}

var ErrUserNotFound = errors.New("user not found")

func findUser(users map[int]string, id int) (string, error) {
	name, ok := users[id]
	
	if !ok {
		return "", ErrUserNotFound
	}

	return name, nil
}

func errorsFile() {

}