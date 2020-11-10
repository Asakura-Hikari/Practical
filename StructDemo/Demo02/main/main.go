package main

import (
	"fmt"
)

/*
使用结构体和方法编写一串代码，使一个3*3的二维数组转置
123			147
456	------> 258
789			369
*/

type TS struct {
	three [3][3]int
}

func (ts *TS) TSmethod() {
	var temp int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > i {
				temp = ts.three[i][j]
				ts.three[i][j] = ts.three[j][i]
				ts.three[j][i] = temp
			}
		}
	}
}

func main() {
	var ts TS
	var num = 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ts.three[i][j] = num
			num++
			fmt.Printf("%d\t", ts.three[i][j])
		}
		fmt.Println()
	}

	fmt.Println()
	ts.TSmethod()

	for _, v := range ts.three {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
