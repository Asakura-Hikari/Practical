package main

import (
	"fmt"
)

/*
定义一个4*4的数组，逐个从键盘输入，并且交换1和4列，2和3列
*/
func main() {

	var change [3][3]int
	r := len(change) - 1

	//输入数字
	for i := 0; i < len(change); i++ {
		for j := 0; j < len(change[i]); j++ {
			fmt.Println("请输入数字: ")
			_, _ = fmt.Scan(&change[i][j])
		}
	}

	//打印
	for i := 0; i < len(change); i++ {
		for j := 0; j < len(change[i]); j++ {
			fmt.Printf("%d\t", change[i][j])
		}
		fmt.Println()
	}

	//交换上下行
	for i := 0; i < r/2; i++ {
		temp := change[i]
		change[i] = change[r]
		change[r] = temp
	}

	fmt.Println()

	//重新打印
	for i := 0; i < len(change); i++ {
		for j := 0; j < len(change[i]); j++ {
			fmt.Printf("%d\t", change[i][j])
		}
		fmt.Println()
	}

}
