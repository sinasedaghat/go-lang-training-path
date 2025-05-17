package main

import "fmt"

func main() {

	prices := []float64{0, 11.11}
	prices = append(prices, 22.22)
	prices = append(prices, 33.33)
	prices = append(prices, 44.44)
	prices = append(prices, 55.55)
	prices = prices[:8]

	fmt.Printf("Original ====>\n\n Array: %v\n Length: %v\n Capacity: %v\n\n\n", prices, len(prices), cap(prices))
	// fmt.Printf("Slice ====>\n\n Array: %v\n Length: %v\n Capacity: %v\n\n\n", slicedPrices, len(slicedPrices), cap(slicedPrices))
	// fmt.Printf("Nested Slice ====>\n\n Array: %v\n Length: %v\n Capacity: %v\n\n\n", nestedSlicePrice, len(nestedSlicePrice), cap(nestedSlicePrice))
}

// func printer(arr []float64) {
// 	fmt.Printf("Array: %v\nLength: %v\nCapacity: %v\n\n\n", arr, len(arr), cap(arr))
// }
