package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
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

func main() {
	Vat, getFund := calculateTax(100)
	fmt.Println(Vat, getFund)

	// Pointers

}
