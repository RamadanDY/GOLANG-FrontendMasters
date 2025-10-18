package main

// global scope
var url = "https://hellojkhgkuyh.com"

func main() {
	// functional scope
	var message string = "hello"
	var price float32 = 34.5
	println(message, price, url)
	println("helo")

	{
		// block scope
	}
}
