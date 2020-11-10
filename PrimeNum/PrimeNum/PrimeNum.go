package PrimeNum

import "fmt"

func PrimeNum(p int) {

	var count, total int = 0,0
	fmt.Printf("0-%d的质数有： \n",p)
	for i := 2; i <= p; i++ {
		a := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				a++
			}
		}

		if a == 0 {
			fmt.Printf("%d ", i)
			total += i
			count++
			if count%5 == 0 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Print("\n")
	fmt.Print("他们的和为: ", total)
}
