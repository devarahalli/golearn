package main

import "fmt"

func greatest(val ...int) int {
	return max(val)

}
func max(val []int) int {
	var res int
	for _, v := range val {
		if res < v {
			res = v
		}
	}
	return res
}

func main() {

	fmt.Println(greatest(100, 10, 20, 30))
}
