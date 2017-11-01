package main

import "fmt"

func main() {
	var res int = 0
	for i := 0; i < 1000; i++ {
		switch {
		//	case (i%3 == 0 && i%5 == 0):
		//res += i
		//fmt.Println(i, "\t\t", "FizzBuzz")

		case i%3 == 0:
			//fmt.Println(i, "\t\t", "Fizz")
			res += i

		case i%5 == 0:
			//fmt.Println(i, "\t\t", "Buzz")
			res += i
		}
	}
	fmt.Println("Total result=", res)
}
