package main

import "fmt"
/*
倒叙打印数组
*/
func main()  {
	odd := []int {1,3,5,7,9}
	r := len(odd) - 1
	temp := 0
	for i:=0;i<r;i++ {
		temp = odd[i]
		odd[i] = odd[r]
		odd[r] = temp
		r--
	}
	fmt.Println(odd)

}
