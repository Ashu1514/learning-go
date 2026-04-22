package main

import "fmt"

func main() {
	age := 32

	// var agePointer *int
	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age Pointer value:", *agePointer)


	fmt.Println(getAdultYears(agePointer)) // pasing pointer means address of value
	// fmt.Println(getAdultYears(&age)) // this also does the same
}

func getAdultYears(age *int) int {
	return *age - 18 // *age to get value at pointer address
}