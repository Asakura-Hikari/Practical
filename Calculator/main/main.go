package main

import (
	c "Practical/Calculator/Calculator"
	"fmt"
)

func  main() {
	var sel int
	for {
		c.Menu()
		_, _ = fmt.Scan(&sel)
		if sel==0 {
			break
		}
		c.Cal(sel)
	}
}
