package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var inCh chan Person
var outCh chan bool

func main() {
	num := runtime.NumCPU()
	fmt.Println("有",num,"个cpu")

	inCh = make(chan Person, 10)
	outCh = make(chan bool, 3)

	//开启十个线程创建数据
	for i := 0; i < 10; i++ {
		go newPerson()
	}

	//第十一个线程用来读取数据
	go readchan()

	//记录所有线程, 当所有线程执行完毕结束主线程
	for i:=0;i<11;i++{
		<-outCh
	}
	close(inCh)
	close(outCh)
	fmt.Println("所有线程执行完毕")
}

func newPerson() {
	p := &Person{"阿卡丽", rand.Intn(100), "艾欧尼亚"}
	inCh <- *p
	outCh <- true
}

func readchan() {
	for i:=0;i<10;i++ {
		v := <-inCh
		fmt.Printf("随机创建的阿卡丽有:%v\n", v)
		println(i+1)
	}
	outCh <- true
}
