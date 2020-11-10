package Bubble

/*
这就是传失踪的冒泡排序
重点：进行内外镶嵌循环 外循环次数为长度-1 内循环次数为长度-外循环的次数
对比内循环中 当前索引 和 当前索引+1的值 并且用临时变量转换
*/

func Bubble(list *[]int) {
	Time := len(*list) - 1
	temp := 0
	for i := 0; i < Time; i++ {
		for j := 0; j < Time-i; j++ {
			if (*list)[j] > (*list)[j+1] {
				temp = (*list)[j]
				(*list)[j] = (*list)[j+1]
				(*list)[j+1] = temp
			}
		}
	}
}
