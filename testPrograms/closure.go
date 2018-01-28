package main

import "fmt"

func TestCounter() func() {

	testVar := 0
	return func() {
		test1Var := 0
		testVar++
		fmt.Println("testVar -->", testVar)
		fmt.Println("testi1Var -->", test1Var)
	}
}

func main() {

	var incr func()
	incr = TestCounter()
	incr()
	incr()
	incr()
	incr()
	fmt.Printf("type of incr %T\n  and %v\n", incr, incr)
}
