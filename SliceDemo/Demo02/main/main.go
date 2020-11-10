package main

import "fmt"

/*
已知有一个已经排好的数组(升序)，要求插入一个元素，最后打印该数组，依然是升序(?)
有问题：右边的数组第一个数字会相等于插入的元素, 重新声明变量解决了，但是原理未知
这个代码使用了二分法
*/

func main() {
	var eve = []int{3, 4, 5, 6, 7, 8, 22, 34, 56, 78,99}
	//array(&eve, 40)
	check(&eve, 0, len(eve)-1, 42)
}

//func array(arr *[]int, val int) {
//	for i := 1; i < len(*arr); i++ {
//
//		if val >= (*arr)[i-1] && val <= (*arr)[i] {
//			fmt.Println(val)
//			left := (*arr)[:i]
//			fmt.Println(left)
//			right := (*arr)[i:]
//			fmt.Println(right)
//			left = append(append(left, val), right...)
//			fmt.Println(left)
//			break
//		} else if (*arr)[i] > val {
//			var left []int
//			left = append(append(left, val), *arr...)
//			fmt.Println(left)
//			break
//		} else if (*arr)[i] < val {
//			var left []int
//			left = append(*arr, val)
//			fmt.Println(left)
//			break
////		}
//	}
//}

func check(arr *[]int, leftindex int, rightindex int, val int) {
	middle := (leftindex + rightindex) / 2

	if leftindex > rightindex {
		fmt.Println("索引结束")
		left := (*arr)[:rightindex+1]
		fmt.Println("左边的数值为", left)
		right := (*arr)[rightindex+1:]
		r1 := right[0]
		l1 := len(left)
		fmt.Println("右边的数值为", right)
		fmt.Println("要插入的值为",val)
		left = append(append(left, val), right...)
		left[l1+1] = r1
		fmt.Println("最终获得", left)
		return
	}

	if (*arr)[middle] > val {
		fmt.Printf("要比较的值在左边，中数为%d, 索引为%d\n", (*arr)[middle], middle)
		check(arr, leftindex, middle-1, val)
	} else if (*arr)[middle] < val {
		fmt.Printf("要比较的值在右边，中数为%d, 索引为%d\n", (*arr)[middle], middle)
		check(arr, middle+1, rightindex, val)
	} else {
		fmt.Println("索引结束")
		left := (*arr)[:middle]
		fmt.Println("左边的数值为", left)
		right := (*arr)[middle:]
		r1 := right[0]
		l1 := len(left)
		fmt.Println("右边的数值为", right)
		fmt.Println("要插入的值为",val)
		left = append(append(left, val), right...)
		left[l1+1] = r1
		fmt.Println("最终获得", left)
	}
}
