package main

import "fmt"

func retFunc() func() string {
	fn := func() string {
		return fmt.Sprintln("Function returning : Hello Func")
	}
	return fn
}

func main() {

	x := retFunc()
	fmt.Println(x())

}
