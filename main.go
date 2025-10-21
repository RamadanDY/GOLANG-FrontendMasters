package main

import (
	"fmt"
)

// global scope
var url = "https://hellojkhgkuyh.com"
var name = "RamDev is here "

func init() {
	fmt.Println("B")
}

// function to calculate the tax
func calculateTax(price float64) (Vat float64, getFund float64) {
	Vat = price * 0.09
	getFund = price * 0.02
	return

}

// the asterisk tell the parameter is a pointer to the int not just any int
func newBD(age *int) {
	fmt.Printf("the pointer is %v and the value is %v", age, *age)
	// we dereference the pointer ie we get access to the value at that pointer
	*age = *age + 1
}
func main() {
	Vat, getFund := calculateTax(100)
	fmt.Println(Vat, getFund)

	// Pointers
	age := 2
	newBD(&age)
	fmt.Println(age)

}
