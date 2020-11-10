package main

func main() {
	var t ten
	{
	}
	var exit chan bool
	exit = make(chan bool, 4)
	tensNum := make(chan int, 20)
	total := make(chan int, 200)

	go t.addNum(total, exit)
	for i := 0; i < 3; i++ {
		go t.coculate(tensNum, total, exit)
	}
	for v:= range exit{
		<-exit
		println(v)
	}
}
