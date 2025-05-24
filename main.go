package main

import "fmt"

func main() {
	price := []float64{12.99, 67.34, 4}
	fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))

	// price = append(price, 21.45, 25.45, 20, 87, 3)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))
	// price = append(price, 21.45, 25.45, 20, 87, 434, 577, 4, 4, 6)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))

	// price = append(price, 21.45)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))
	// price = append(price, 25.45)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))
	// price = append(price, 20)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))
	// price = append(price, 20.09)
	// fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))

	newPrice := []float64{254, 98.34, 99, 99}
	fmt.Printf("newPrice: %v, length: %v, capacity: %v\n", newPrice, len(newPrice), cap(newPrice))

	price = append(price, newPrice...)
	fmt.Printf("price: %v, length: %v, capacity: %v\n", price, len(price), cap(price))

}
