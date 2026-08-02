package main

import "fmt"

func main() {
	age := 32 // Biến thông thường

	agePointer := &age

	fmt.Println("Age (Pointer Address):", agePointer)
	fmt.Println("Age (Value from Pointer):", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Age (After Mutation):", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
