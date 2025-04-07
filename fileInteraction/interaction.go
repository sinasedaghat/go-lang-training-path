package fileInteraction

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteNumberToFile(fileName string, balance float64) {
	os.WriteFile(fileName, []byte(fmt.Sprint(balance)), 0644)
}

func GetNumberFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	// fmt.Println("🔮 This is ERROR received when i want get balance data from file.", err)
	if err != nil {
		return 0, errors.New("your Balance data does not exists")
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	// fmt.Println("🔮 This is ERROR received when i want convert string from file to floating point number.", err)
	if err != nil {
		return 0, errors.New("your Balance isn't valid value")
	}
	return balance, nil
}
