package main

import (
	"fmt"
)

/*
定义一个三行四列的二维数组，逐个从键盘输入，并且将四周清零
*/

func main() {
	var zero [5][5]int

	//逐个输入数字
	for i := 0; i < len(zero); i++ {
		for j := 0; j < len(zero[i]); j++ {
			fmt.Println("请输入数字: ")
			_, _ = fmt.Scan(&zero[i][j])
		}
	}

	//将数字四周清零
	for i := 0; i < len(zero); i++ {
		for j := 0; j < len(zero[i]); j++ {
			if i == 0 || i == len(zero)-1 {
				zero[i][j] = 0
			}
			if j == 0 || j == len(zero[i])-1 {
				zero[i][j] = 0
			}
		}
	}

	//输出数字
	for i := 0; i < len(zero); i++ {
		for j := 0; j < len(zero[i]); j++ {
			fmt.Printf("%d\t",zero[i][j])
		}
		fmt.Println()
	}
}
