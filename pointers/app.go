package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int
	agePointer := &age // pointer variable

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age Pointer value:", *agePointer)


	// fmt.Println(getAdultYears(agePointer)) // pasing pointer means address of value
	// fmt.Println(getAdultYears(&age)) // this also does the same as above func calling

	editAgeToAdultYears(agePointer) // manipulated value in function and stored at same address agian using pointers
	fmt.Println("Adult Age:", age) // printing original variable
}

func editAgeToAdultYears(age *int) {
	// return *age - 18 // *age to get value at pointer address
	*age = *age - 18 // manipulating value at pointer and restore it at same pointer 
}