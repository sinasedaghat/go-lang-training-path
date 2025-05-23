package products

import "fmt"

type Product struct {
	id    int
	Title string
	Price float64
}

var Products = []Product{
	{
		id:    0,
		Title: "Mobile",
		Price: 199000.99,
	},
	{
		id:    1,
		Title: "Laptop",
		Price: 25000.00,
	},
}

func NewProduct(title string, price float64) *Product {
	return &Product{
		id:    len(Products) + 1,
		Title: title,
		Price: price,
	}
}

func AddProduct(product *Product) {
	Products = append(Products, *product)
}

func ListOfProducts() {
	fmt.Println("List of product: ", Products)
}

func Yo() {
	fmt.Println("Product package")
}
