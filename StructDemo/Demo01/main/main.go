package main

import "fmt"
/*
调用结构体，给出结构体方法并且调用该方法 - 打赢输出数字的对应乘法表
*/
type MethodUtils struct {
	//Nothing
}

func (MT *MethodUtils) MultiTab(num int) {
	for i := 1; i <= num; i++ {
		for j:=1; j<=i;j++{
				fmt.Printf("%d x %d = %d\t",j,i,i*j)
		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	var i int
	fmt.Println("输入你想看的乘法表:")
	_, _ = fmt.Scanln(&i)
	mu.MultiTab(i)
}
