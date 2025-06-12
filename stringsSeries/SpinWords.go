package stringsSeries

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/5264d2b162488dc400000001/go

func SpinWords(str string) string {
	var result []string

	for _, word := range strings.Split(str, " ") {
		// for word := range strings.SplitSeq(str, " ") {
		if len(word) >= 5 {
			word = inverse(word)
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}

func inverse(str string) string {
	var invertedString string
	for i := len(str) - 1; i >= 0; i-- {
		invertedString += string(str[i])
	}
	return invertedString
}

// another solution
func AnotherSpinWords(str string) string {
	list := strings.Split(str, " ")

	for index, word := range list {
		if len(word) >= 5 {
			list[index] = anotherInverse(word)
		}
	}
	return strings.Join(list, " ")
}

func anotherInverse(str string) string {
	list := []rune(str)
	fmt.Println("before", list)

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	fmt.Println("after", list)
	return string(list)
}
