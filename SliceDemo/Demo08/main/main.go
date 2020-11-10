package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成一个数组，数组有五个数，找出其最大值和最小值以及下标
*/
func main(){
	rand.Seed(time.Now().UnixNano())
	var arr[5] int
	large := 0
	larIndex,SmaIndex :=0,0
	small := 100

	for i:=0; i<5; i++{

		//创建随机数并且加入数组
		arr[i] = rand.Intn(100)

		//查找最大数
		if arr[i] > large {
			large = arr[i]
			larIndex = i
		}

		//查找最小数
		if arr[i] < small{
			small = arr[i]
			SmaIndex = i
		}
	}

	//输出
	fmt.Println(arr)
	fmt.Printf("最大数为%d, 下标为%d\n", large,larIndex)
	fmt.Printf("最小数为%d, 下标为%d", small,SmaIndex)

}
