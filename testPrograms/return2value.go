package main

import "fmt"
import "flag"

func twoReturn(s int) (int, bool) {
	var x bool

	if s%2 == 0 {
		x = true
	} else {
		x = false
	}
	r := s / 2
	return r, x

}

func main() {
	v := flag.Int("i", 0, "Enter integer value")
	flag.Parse()
	x, y := twoReturn(*v)
	fmt.Println(x, y)
}
