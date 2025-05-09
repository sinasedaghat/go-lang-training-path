package main

import "fmt"

type addInput interface {
	int | float64 | string
}

func main() {
	// person := struct {
	// 	Age int
	// }{
	// 	Age: 4,
	// }
	// printAnything(4)
	// printAnything(4.9)
	// printAnything("4")
	// printAnything(person)

	// printSomething(4)
	// printSomething(4.9)
	// printSomething("4")
	// printSomething(person)

	primevalSum := primevalAdd(1, 6)
	fmt.Printf("primevalSum: %v\nprimevalSum type: %T\n", primevalSum, primevalSum)
	// primevalSum += 1 // error: mismatched types any and int

	sum := add(1, 6)
	fmt.Printf("sum: %v\nsum type: %T\n", sum, sum)
	sum += 1 // its Okay

}

func printAnything(input interface{}) {
	// switch inputType := input.(type) {
	switch input.(type) {
	case int:
		{
			fmt.Printf("This input is Int\ninput=%v\n", input)
		}
	case float64:
		{
			fmt.Printf("This input is Float64\ninput=%v\n", input)
		}
	case string:
		{
			fmt.Printf("This input is String\ninput=%v\n", input)
		}
	default:
		{
			fmt.Printf("input=%v\n", input)
		}
	}
}

func printSomething(value any) {
	intVal, isOK := value.(int)
	if isOK {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, isOK := value.(float64)
	if isOK {
		fmt.Println("Float: ", floatVal)
		return
	}

	strVal, isOK := value.(string)
	if isOK {
		fmt.Println("String: ", strVal)
		return
	}

	fmt.Println("Any: ", value)
}

func primevalAdd(a, b any) any {
	aInt, aIsOK := a.(int)
	bInt, bIsOK := b.(int)

	if aIsOK && bIsOK {
		return aInt + bInt
	}

	aFloat, aIsOK := a.(float64)
	bFloat, bIsOK := b.(float64)

	if aIsOK && bIsOK {
		return aFloat + bFloat
	}

	aString, aIsOK := a.(string)
	bString, bIsOK := b.(string)

	if aIsOK && bIsOK {
		return aString + bString
	}

	return nil

}

func add[T addInput](a, b T) T {
	return a + b
}
