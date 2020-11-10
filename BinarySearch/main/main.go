package main

import (
	"Practical/BinarySearch/BS"
	"fmt"
)

func main()  {
	var i int
	fmt.Println("输入你想查找的值：")
	fmt.Scanln(&i)
	db := []int {1,3,5,7,9,11,13,15,17}
	BS.BS(&db,0,len(db)-1,i)

}
