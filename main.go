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

func newBD(age *int) {
	fmt.Printf("the pointer is %v and the value is %v", age, *age)
	*age++
}
func main() {
	Vat, getFund := calculateTax(100)
	fmt.Println(Vat, getFund)

	// Pointers
	age := 22
	newBD(&age)
	fmt.Println(age)

}
