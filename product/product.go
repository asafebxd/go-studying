package product

import "fmt"

type Product struct {
	name string
	price float64
}

func New(name string, price float64) Product {
	return Product{
		name: name,
		price: price,
	}
}

func (product Product) Display() string {
	return fmt.Sprintf("%s costs %.2f", product.name, product.price)
}