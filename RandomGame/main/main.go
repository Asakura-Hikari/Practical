package main

import (
	r "Practical/RandomGame/RandomGame"
	"fmt"
)

func main(){
	var gue, count, chance int
	ran := r.GetRandom()

	for i:=0;i<10;i++{
		chance = 10-count
		fmt.Printf("你还有%d次机会!",chance)
		fmt.Println("请输入你猜测的数字: ")
		count++
		_, _ = fmt.Scanln(&gue)
		if !r.Match(ran,gue) {
			continue
		}else if count == 1{
			println("你真是个天才！")
			break
		}else if count < 4 {
			println("你很聪明哦")
			break
		}else if count < 10 {
			println("马马虎虎啦")
			break
		}else if count == 10{
			println("这可是最后一次机会, 你真走运！")
		}
	}
	println("游戏结束")


}
