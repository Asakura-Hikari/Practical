package main

import (
	"Practical/BubbleSort/Bubble"
	"fmt"
)

func main() {
	arr := []int{43, 23, 65, 34, 68, 25}
	Bubble.Bubble(&arr)
	fmt.Println(arr)
}
