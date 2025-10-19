package main

// global scope
var url = "https://hellojkhgkuyh.com"
var name = "RamDev is here "

func main() {
	// functional scope
	var message string = "hello"
	var price float32 = 34.5
	println(message, price, url)
	println("helo")
	printData()
	{
		// block scope
	}
}
