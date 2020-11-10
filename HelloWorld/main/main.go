package main

import "fmt"

func main() {
	var left = []int{3, 4, 5, 6, 7, 8}
	var right = []int{22,34,56}
	var v int = 9

	left = append(append(left,v),right...)
	fmt.Println(left)
}
