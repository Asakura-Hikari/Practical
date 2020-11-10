package RandomGame

import (
	"fmt"
	"math/rand"
	"time"
)

//获取一个随机数字
func GetRandom() int{
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(101)
	return r
}

//对比数字并且提示
func Match(ran, gue int) bool{
	a:=false
	if gue == ran {
		a = true
	}else if gue > ran{
		fmt.Println("你猜的数字太大")
	}else {
		fmt.Println("你猜的数字太小")
	}
	return a
}


