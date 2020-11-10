package main

import (
	"Practical/OrderFind/OrderFind"
	"fmt"
)

func main(){
	var hero string
	fmt.Println("输入你想查找的英雄")
	fmt.Scanln(&hero)
	o := OrderFind.OF(hero)
	if o {
		fmt.Println("这个英雄存在")
	}else{
		fmt.Println("英雄不存在")
	}

}
