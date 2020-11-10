package main

import (
	"Practical/Rectangle/Rectengle"
	"fmt"
)

func main() {
	var a int
	fmt.Println("Please enter size:")
	fmt.Scan(&a)
	Rectengle.Rectangle(a)
}
