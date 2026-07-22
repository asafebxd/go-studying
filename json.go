package main

import (
	"encoding/json"
	"os"
)

type Produto struct {
	Name string `json:"name"`
	Price float64  `json:"price"`
	Stock int `json:"stock"`
}

func productToJSON(product Produto) (string, error) {
	data, err := json.Marshal(product)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func productFromJSON(data string) (Produto, error) {
	var produto Produto

	err := json.Unmarshal([]byte(data), &produto)
	if err != nil {
		return Produto{}, err
	}

	return produto, nil
}

func saveProducts(filename string, products []Produto) error {
	data, err := json.MarshalIndent(products, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
} 

func loadProducts(filename string) ([]Produto, error) {
	data, err := os.ReadFile(filename) 
	if err != nil {
		return nil, err
	}

	var products []Produto

	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func jsonLesson() {
	
}