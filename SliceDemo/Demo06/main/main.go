package main

import "fmt"

/*
实现查找核心代码。
比如已知有个arr [10]String的数组 查找"AA" 是否存在 并且给出下标（多个AA给出多个下标）
*/

func main(){

	arr := []string {"AA","BB","CC","DD","EE","AA","BB","CC","DD","AA"}
	var in string
	b := false
	slice := make([]int,0)

	fmt.Println("输入你要查找的字符: ")
	_, _ = fmt.Scanln(&in)

	//查找字符，并且加入切片
	for i:=0;i<len(arr);i++ {
		if in == arr[i]{
			b = true
			slice = append(slice,i)
		}
	}

	if b {
		fmt.Println("数据库中有该字符，下标为: ",slice)
	} else {
		fmt.Println("数据库中没有该字符")
	}

}
