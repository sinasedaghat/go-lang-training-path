package cmdmanager

import "fmt"

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (cmdm CMDManager) InputReader() ([]string, error) {
	fmt.Println("please enter your prices")
	fmt.Println("after enter all prices type 'exit'")
	prices := make([]string, 0)

	for {
		price := ""
		fmt.Print("Enter your price: ")
		fmt.Scan(&price)
		if price == "exit" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmdm CMDManager) OutputWriter(data any) error {

	fmt.Println(data)
	return nil
}
