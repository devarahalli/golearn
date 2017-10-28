package main

import "fmt"

const (
	q  string  = "This is cool"
	Pi float32 = 3.14
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {

	const p = 42
	fmt.Println("q --> ", q)
	fmt.Println("p --> ", p)
	fmt.Println("p --> ", Pi)

	fmt.Printf("Iota of KB Binary %b\n", KB)
	fmt.Printf("Iota of KB Decimal %d\n", KB)
	fmt.Printf("Iota of MB Binary %b\n", MB)
	fmt.Printf("Iota of MB Decimal %d\n", MB)

	var x float64 = 23
	res := x + Pi
	fmt.Println(res)
}

//fmt.Println(q)
