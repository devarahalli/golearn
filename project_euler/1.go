package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if multipleofthreeorfive(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("Sum multiple of 3 or 5", sum)
}

func multipleofthreeorfive(nn int) bool {
	if nn%3 == 0 {
		return true
	}
	if nn%5 == 0 {
		return true
	}
	return false

}
