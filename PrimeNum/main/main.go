package main

import (
	p "Practical/PrimeNum/PrimeNum"
	"fmt"
)

func main(){
	var pri int
	fmt.Println("输入你想知道的数值范围:")
	fmt.Scan(&pri)
	p.PrimeNum(pri)
}
