package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)

// global scope
var url = "https://hellojkhgkuyh.com"
var name = "RamDev is here "

func main() {
	// functional scope
	var message string = "hello"
	var price float32 = 34.5
	fmt.Println(message, price, url, name, data.Name)
	fmt.Println("helo")

	SendMessage()
	{
		// block scope
	}
}
