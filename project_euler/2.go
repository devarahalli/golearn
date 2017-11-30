package main

import "fmt"

func fib() func() int {
	nn := 1
	res := 1
	temp := 0
	return func() int {
		//	fmt.Println(nn, res)
		temp = nn
		nn = res + nn
		res = temp
		return nn
	}
}

func main() {
	sum := 0
	fn := fib()
	for true {
		y := fn()
		if y > 4000000 {
			break
		}
		if y%2 == 0 {
			sum += y

		}
	}
	fmt.Println(sum)
}
