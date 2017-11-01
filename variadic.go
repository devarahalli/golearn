package main

import "fmt"

func average(num ...float64) float64 {
	fmt.Printf("type of num is =%T\n", num)
	fmt.Println(num)
	var total float64
	for _, v := range num {
		total += v
	}
	return total / float64(len(num))
}

func main() {
	fmt.Println("This program calculates the average of intetger number")
	//res := average(1, 2, 3, 4, 5, 6, 7)
	res := average(1, 2, 3)
	fmt.Println("Result is =", res)

}
