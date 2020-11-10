package main

import (
	m "Practical/Monthday/monthday"
	"fmt"
)

func main() {

	for {
		var month, year, day int
		var control string
		fmt.Println("输入你想知道的年份: ")
		_, _ = fmt.Scanln(&year)
		fmt.Println("输入你想知道的月份: ")
		_, _ = fmt.Scanln(&month)
		day = m.Monthday(year, month)
		fmt.Printf("%d年%d月有 : %d 天\n", year, month, day)
		fmt.Println("输入 C 继续, 输入任意键结束。")
		_, _ = fmt.Scanln(&control)

		if control == "C"{
			continue
		}else{
			break
		}

	}
}
