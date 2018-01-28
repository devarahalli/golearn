package main

import "fmt"

func average(num ...float64) float64 {
	fmt.Printf("type of num is =%T\n", num)
	var total float64
	for _, v := range num {
		total += v
	}
	return total / float64(len(num))
}

func main() {
	data := [7]float64{1, 2, 3, 4, 5, 6, 7}
	//	res := data[:]
	//	fmt.Printf("Inside main func type of res = %T\n", res)
	rt := average(data[:]...)
	fmt.Println("Result is =", rt)
}
