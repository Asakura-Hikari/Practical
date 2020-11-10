package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
定义一个8个数字的数组，并且求出大于平均数的个数，和小于平均数的个数
*/
func  main()  {

	rand.Seed(time.Now().UnixNano())
	var total, bave, lave int
	var arr[8] int

	//随机生成8个数并且加入数组
	for i:=0; i<8; i++{
		arr[i] = rand.Intn(100)
		total +=arr[i]
	}

	//计算平均数
	ave := total/len(arr)

	//大于和小于平均数数字的个数
	for _, v := range arr{
		if v > ave {
			bave++
		}
		if v < ave {
			lave++
		}
	}

	fmt.Println(arr)
	fmt.Printf("平均数为%d\n",ave)
	fmt.Printf("有%d数大于平均数\n",bave)
	fmt.Printf("有%d数小于平均数\n",lave)

}