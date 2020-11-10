package Calculator

import (
	"fmt"
)

func Sum(i, j float32) float32 {
	return i + j
}

func Sub(i, j float32) float32 {
	return i - j
}

func Mul(i, j float32) float32 {
	return i * j
}

func Div(i, j float32) float32 {
	return i / j
}

func Menu() {
	fmt.Println("----------------小小计算器----------------")
	fmt.Println("1.加法 ")
	fmt.Println("2.减法 ")
	fmt.Println("3.乘法 ")
	fmt.Println("4.除法 ")
	fmt.Println("0.退出 ")
	fmt.Println("----------------------------------------")

}

func Cal(i int) {

	var f, s float32
	if i > 4 || i < 0 {
		fmt.Println("输入不正确，请重新输入")
	} else {

		fmt.Println("请输入第一个数字: ")
		_, _ = fmt.Scan(&f)
		fmt.Println("请输入第二个数字: ")
		_, _ = fmt.Scan(&s)

		switch i {
		case 1:
			fmt.Printf("%.2f+%.2f=%.2f\n", f, s, Sum(f, s))
		case 2:
			fmt.Printf("%.2f+%.2f=%.2f\n", f, s, Sub(f, s))
		case 3:
			fmt.Printf("%.2f+%.2f=%.2f\n", f, s, Mul(f, s))
		case 4:
			fmt.Printf("%.2f+%.2f=%.2f\n", f, s, Div(f, s))
		}
	}
}
