package main

import (
	"Practical/FN/FN"
	"fmt"
)

func main() {
	var year, month, day int

	for {
		fmt.Println("请输入年份: ")
		_, _ = fmt.Scanln(&year)
		if year < 0 || year > 5000 {
			fmt.Println("输入的年份有误")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("请输入月份: ")
		_, _ = fmt.Scanln(&month)
		if month < 0 || month > 12 {
			fmt.Println("输入的月份有误")
			continue
		} else {
			break
		}

	}

	for {
		fmt.Println("请输入几号: ")
		_, _ = fmt.Scanln(&day)
		if day < 0 || day > 31 {
			fmt.Println("输入的日期有误")
			continue
		} else {
			break
		}

		FN.CompareTwoTimes(year, month, day)
	}
}

