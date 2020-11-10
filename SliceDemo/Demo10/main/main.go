package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
跳水比赛，8个评委打分，得到的成绩是去掉一个最高分，去掉一个最低风，剩下6个数的平均分，使用一维数组实现：
1. 找到最高分评委和最低分评委
2. 找到最佳评委(和最后得分最相近)和最差评委（和最后得分相差最远）
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	var mark [8]float64
	var ab [len(mark)]float64

	lar := 0.00
	sma := 100.00
	total := 0.00
	lab :=0.00
	sab :=100.00
	lindex := 0
	sindex := 0

	//随机给出分数并且加入数组
	for i := 0; i < len(mark); i++ {
		mark[i] = float64(rand.Intn(100)) + rand.Float64() + 1
		fmt.Printf("裁判%d给的数是: %.2f \n", i+1, mark[i])
		if mark[i] > lar {
			lar = mark[i]
		}
		if mark[i] < sma {
			sma = mark[i]
		}
	}
	fmt.Println()

	//找到最高分和最低分的评委并且排除
	for i := 0; i < len(mark); i++ {
		if mark[i] == lar {
			fmt.Printf("裁判%d给的分最高，分数为%.2f, 已去除\n", i+1, mark[i])
			mark[i] = 0.00
		}
		if mark[i] == sma {
			fmt.Printf("裁判%d给的分最低，分数为%.2f, 已去除\n", i+1, mark[i])
			mark[i] = 0.00
		}
		total += mark[i]

	}

	//计算出平均分
	ave := total / float64(len(mark)-2)
	fmt.Printf("平均分为%.2f\n", ave)

	//求出排除之后的裁判与平均数的差并且加入数组ab
	for i := 0; i < len(mark); i++ {
		if mark[i] != 0.00 {
			ab[i] = math.Abs(mark[i] - ave)
		}else {
			ab[i] = 0
		}
	}

	//求出排除之后的裁判的具体索引位置和具体数值
	for i:=0; i<len(ab); i++{
		if ab[i] > lab && ab[i] != 0{
			lab = ab[i]
			lindex = i
		}
		if ab[i] < sab && ab[i] != 0{
			sab = ab[i]
			sindex = i
		}
	}

	fmt.Printf("最佳评委为%d号评委, 与平均分相差%.2f\n",sindex+1,sab)
	fmt.Printf("最差评委为%d号评委, 与平均分相差%.2f\n",lindex+1,lab)

}
