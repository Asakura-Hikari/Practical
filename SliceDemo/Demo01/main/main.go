package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成10个整数，保存到数组，并且倒叙，求平均值，最大值，和最大值到下标，并且检测里面是否有55
*/

func main() {

	//声明变量
	rand.Seed(time.Now().UnixNano())
	var random[10] int
	var total, temp, index int
	var a bool
	large := random[0]
	rightindex := len(random)-1

	//把随机数加入数组
	for i := 0; i<len(random); i++ {
		random[i] = rand.Intn(100)
		fmt.Printf("%d ", random[i])
		total += random[i]
	}

	fmt.Println()
	fmt.Printf("平均值是%d \n倒叙为: ",total/len(random))

	//求出最大值和索引值
	for i, v:= range random{
		if large < v {
			large = v
			index = i
		}
		if v==55 {
			a = true
		}
	}

	//倒叙
	for i:=0;i<len(random)/2;i++ {
		temp = random[i]
		random[i] = random[rightindex]
		random[rightindex] = temp
		rightindex--
	}

	//打印倒叙
	for _, v := range random{
		fmt.Printf("%d ",v)
	}

	//输出最终结果
	if a == true {
		fmt.Printf("\n最大值为%d,下标为%d, 数组中含有55", large,index)
	}else{
		fmt.Printf("\n最大值为%d,下标为%d, 数组中没有55", large,index)
	}
}
