package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("age from copy variable in function (use normal variable) ==> ", age)

	fmt.Println("age from variable (use variable contain pointer) ==> ", *agePointer)

	fmt.Println("this is variable contain pointer of age ==> ", agePointer)

	fmt.Println("this is pointer of age variable ==> ", &age)

	fmt.Printf("this age before call changeValueReference(&age as *int == %v) ===> %v\n", &age, age)
	changeValueReference(&age)
	fmt.Printf("this age after call changeValueReference(&age as *int == %v) ===> %v\n", &age, age)
}

func changeValueReference(age *int) {
	fmt.Println("this argument of changeValueReference() function", age)
	fmt.Println("before change age in changeValueReference() function", *age)

	*age = 10

	fmt.Println("after change age in changeValueReference() function", *age)
}
