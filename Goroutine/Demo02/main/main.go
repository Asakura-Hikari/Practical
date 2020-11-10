package main

import (
	"fmt"
	"runtime"
)
/*
开启一个协程创建2000个数字，再开启8个协程（线程），计算1-2000的每一个值，再开启一个协程读取
*/
func main() {
	num := runtime.NumCPU()
	fmt.Printf("有%v个cpu,同时开始执行任务\n", num)

	var numChan chan int
	var resChan chan int
	var exitChan chan bool

	numChan = make(chan int, 200)
	resChan = make(chan int, 200)
	exitChan = make(chan bool, 10)

	go Creater(numChan, exitChan)

	for i := 0; i<8;i++ {
		go Coculator(numChan,resChan,exitChan)
	}

	go Reader(resChan,exitChan)

	for i := 0; i < 10; i++ {
		<-exitChan
	}

	fmt.Println("所有线程执行完毕，程序结束")
	close(numChan)
	close(resChan)
	close(exitChan)
}

func Creater(numChan chan int, exitChan chan bool) {
	for i := 1; i <= 200; i++ {
		numChan <- i
	}
	fmt.Println("创建完毕")
	exitChan <- true
}

func Coculator(numChan chan int, resChan chan int, exitChan chan bool) {
	fmt.Println("线程开启！")

forlevel:
	for {
		select {
		case n := <-numChan:
			a := 0
			for i:=0; i<n; i++ {
				a += i
			}
			resChan <- a
		default:
			break forlevel
		}
	}
	exitChan <- true
	fmt.Println("线程结束")
}

func Reader(resChan chan int, exitChan chan bool) {
	for i :=0;i<200;i++{
		res := <-resChan
		fmt.Printf("计算出的结果为：%d\n",res)
	}
	exitChan <- true
}
