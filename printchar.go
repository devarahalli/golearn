package main

import "fmt"
import "time"

func main() {

	for i := 50; i <= 10*1000; i++ {
		fmt.Printf("%b - %v - %v - %v\n", i, i, string(i), []byte(string(i)))
		time.Sleep(10 * time.Millisecond)
	}
}
