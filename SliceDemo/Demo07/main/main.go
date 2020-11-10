package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成10个整数(1-100),使用冒泡排序将其排序，并且用二分法查找是否有90这个数字并且显示下标
*/

func main()  {
	//创建随机种子
	rand.Seed(time.Now().UnixNano())
	Arr := make([]int,0)

	//随机生成10个数字并且加入数组
	for i:=0;i<10;i++{
		r := rand.Intn(100)+1
		Arr = append(Arr,r)
	}
	fmt.Println(Arr)
	fmt.Println()

	//对其进行冒泡排序
	l := len(Arr)-1
	temp := 0
	for i:=0;i<l;i++{
		for j :=0; j<l-i;j++{
			if Arr[j] > Arr[j+1] {
				temp = Arr[j]
				Arr[j] = Arr[j+1]
				Arr[j+1] = temp
			}
		}
	}
	fmt.Println(Arr)
	fmt.Println()

	//调用二分法
	BinaryS(&Arr,0,len(Arr)-1,90)
}

//二分法查找实现
func BinaryS (arr *[]int,left int, right int, val int){
	middle := (left + right)/2

	if left > right{
		fmt.Println("没找到该数据")
		return
	}

	if (*arr)[middle] > val{
		fmt.Printf("要比较的值在左边，中数为%d, 索引为%d\n", (*arr)[middle], middle)
		BinaryS(arr, left,middle-1,val)
	}else if (*arr)[middle] < val{
		fmt.Printf("要比较的值在右边，中数为%d, 索引为%d\n", (*arr)[middle], middle)
		BinaryS(arr, middle+1,right,val)
	}else {
		fmt.Printf("数据库中有该数据，数据下标为%d",middle)
	}

}