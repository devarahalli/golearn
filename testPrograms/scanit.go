package main

import "fmt"

func main() {
	var first, last int

	fmt.Print("Enter big num: ")
	fmt.Scanln(&first)
	fmt.Print("Enter small num: ")
	fmt.Scanln(&last)
	fmt.Println(first / last)
	fmt.Println("Your name is ", first, last, "check for new line")
}
