package main

type ten struct {
}

func (t ten) addNum(total chan int,exit chan bool) chan int {
	for i := 0; i < cap(total); i++ {
		total <- i + 1
	}
	close(total)
	exit <- true
	return total
}

func (t ten) coculate(tensNum chan int, total chan int, exit chan bool) chan int {
	for v := range total{
		if v%10 == 0{
			tensNum <- v
		}
	}
	exit <- true
	return tensNum
}

func(t ten)printer(tensNum chan int, exit chan bool){

}