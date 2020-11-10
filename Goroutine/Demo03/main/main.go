package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
开启一个协程writeDataToFile 随机生成1000个数字，并且写入文档(完成)
当写完之后 创建sort协程对其进行冒泡排序 然后写入另外一个文档
*/

func main() {
	var exitChan chan bool
	var contect chan bool
	exitChan = make(chan bool, 2)
	contect = make(chan bool, 10)
	go writeDataToFile(exitChan, contect)
	go sort(exitChan, contect)
	for i:=0;i<2;i++{
		<-exitChan
	}
	close(exitChan)
}

//创建一个文档并且写入1000个随机数
func writeDataToFile(exitChan chan bool, contect chan bool) {
	rand.Seed(time.Now().UnixNano())
	FilePath := "/Users/ace/Desktop/RandomData.txt"
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		println("创建文件失败, err =", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		r := rand.Intn(1001)
		sr := fmt.Sprint(r)
		_, err = writer.WriteString(sr + " ")
		if err != nil {
			println("写入文件失败，err =", err)
		}
	}
	_ = writer.Flush()
	exitChan <- true
	contect <- true
	close(contect)
}

//读取文件并且写入数组
func sort(exitChan chan bool, contect chan bool) {
	<-contect
	//打开读取文件
	FilePath := "/Users/ace/Desktop/RandomData.txt"
	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0777)
	if err != nil {
		println("打开文件失败, err =", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var ranlist []int
	ranlist = make([]int, 0)

	//进行数据读取并且储存在数组上
	for i := 0; i < 1000; i++ {
		sr, err := reader.ReadString(' ')
		if err == io.EOF {
			fmt.Printf("文件读取完毕")
		} else if err != nil {
			fmt.Printf("读取文件失败, err = %v\n", err)
		}
		sr = strings.ReplaceAll(sr, " ", "")
		sri, err := strconv.Atoi(sr)
		if err != nil {
			fmt.Printf("转换失败, err = %v\n", err)
		}
		ranlist = append(ranlist, sri)
	}

	fmt.Println()
	//对数组进行冒泡排序
	for i := 0; i < len(ranlist)-1; i++ {
		for j := 0; j < len(ranlist)-i-1; j++ {
			if ranlist[j] > ranlist[j+1] {
				temp := ranlist[j]
				ranlist[j] = ranlist[j+1]
				ranlist[j+1] = temp
			}
		}
	}

	FilePath2 := "/Users/ace/Desktop/BubbleData.txt"
	file2, err := os.OpenFile(FilePath2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		println("创建文件失败, err =", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file2)

	for i := 0; i < len(ranlist); i++ {
		b := ranlist[i]
		sb := fmt.Sprint(b)
		_, err = writer.WriteString(sb + " ")
		if err != nil {
			println("写入文件失败，err =", err)
		}
	}
	_ = writer.Flush()
	exitChan <- true
}
