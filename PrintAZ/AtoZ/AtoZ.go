package AtoZ

import "fmt"

func AtoZ(){
	var lower byte = 97
	for i:=0;i<26;i++{
		fmt.Printf("%c ", lower)
		lower++
	}
	fmt.Println()
	var upper byte = 90
	for i:=0;i<26;i++{
		fmt.Printf("%c ", upper)
		upper--
	}
}
