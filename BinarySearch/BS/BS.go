package BS

import "fmt"

/*
	使用递归写法的二分查找
*/

func BS(array *[]int, leftindex int, rightindex int, findval int){
	middle := (leftindex + rightindex)/2

		//没有找到
	if leftindex > rightindex {
		fmt.Println("数据库中没有该数据")
		return
	}
		//从左边开始找
	if (*array)[middle] > findval {
		 BS(array, leftindex, middle-1, findval )

		 //从右边开始找
	}else if(*array)[middle] < findval{
		BS(array, middle+1, rightindex, findval )

		//找到了
	}else{
		fmt.Printf("数据库中拥有该数据， 下标为%d \n" , middle)
	}
}
