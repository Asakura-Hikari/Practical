package Rectengle
import "fmt"

/*
	掏空矩形的函数，自己写着玩儿的。
 */

func Rectangle(level  int){

	for i := 0; i < level; i++ {
		for j := 0; j < level-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			if j == 0 {
				fmt.Print("*")
			} else if j == i*2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	for i := 0; i < level-1; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<(level - 1) *2 -i*2 -1; j++{
			if j==0 {
				fmt.Print("*")
			}else if j== (level - 1) *2 -i*2 - 2{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
