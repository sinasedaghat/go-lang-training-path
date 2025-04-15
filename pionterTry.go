package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("age from copy variable in function (use normal variable) ==> ", age)

	fmt.Println("age from variable (use variable contain pointer) ==> ", *agePointer)

	fmt.Println("this is variable contain pointer of age ==> ", agePointer)

	fmt.Println("this is pointer of age variable ==> ", &age)
}
