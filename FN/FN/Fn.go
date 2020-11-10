package FN

import "fmt"

func GetTotalDay(year, month, day int) int {
	var total int

	//计算年份的总日期
	for i := 0; i < year; i++ {
		total += 31*7 + 30*4
		if year%4 == 0 && year%100 != 0 {
			total += 29
		} else {
			total += 28
		}
	}

	//计算月份的总日期
	for i := 0; i < month; i++ {
		switch i {
		case 1:
			total += 31
		case 2:
			if year%4 == 0 && year%100 != 0 {
				total += 29
			} else {
				total += 28
			}
		case 3:
			total = +31
		case 4:
			total = +30
		case 5:
			total = +31
		case 6:
			total = +30
		case 7:
			total = +31
		case 8:
			total = +31
		case 9:
			total = +30
		case 10:
			total = +31
		case 11:
			total = +30
		}
	}

	//加上几号
	total += day
	return total
}

func CompareTwoTimes(year, month, day int) {
	var sub = GetTotalDay(year, month, day) - GetTotalDay(1990, 1, 1)
	//println(sub)
	if sub%5 > 2 {
		fmt.Printf("%d年%d月%d日在晒网!", year, month, day)
	} else if sub%5 <= 2 {
		fmt.Printf("%d年%d月%d日在打鱼!", year, month, day)
	}
}
