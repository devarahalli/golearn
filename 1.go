package main

import "fmt"

func hello() func() {
	x := 0
	return func() {
		x += 1
		fmt.Printf("Hello --->%d\n", x)
	}
}

func world2() {
	fmt.Print("World ---> 2\n")
}
func world1() {
	fmt.Print("World ---> 1\n")
}
func checkdefer() {
	fmt.Print("THIS IS INTRESTING  -->\n")
}

func main() {
	defer world2()
	defer checkdefer()
	fn := hello()
	fn()
	fn()
	fn()
	fn()
	world1()
	defer fn()
	fn()

}
