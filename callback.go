package main

import "fmt"

func testcallback(num []int, callback func(int) bool) []int {
	rt := []int{}
	for _, v := range num {
		if callback(v) {
			rt = append(rt, v)
		}
	}
	return rt
}

func main() {

	rv := testcallback([]int{1, 2, 3, 4, 5}, func(num int) bool {
		return num > 1
	})
	fmt.Println(rv)

}
