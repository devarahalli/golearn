package main

import "fmt"
import "flag"

func moduleInt(v int) func() (int, bool) {

	return func() (int, bool) {
		var t bool
		val := v / 2

		if v%2 == 0 {
			t = true
		} else {
			t = false
		}
		return val, t
	}
}

func main() {
	val := flag.Int("val", 20, "User input value")
	flag.Parse()
	fn := moduleInt(*val)
	fmt.Println(fn())

}
