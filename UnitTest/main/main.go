package main

import (
	"Practical/UnitTest/Unit"
)

func main()  {

	monster := Unit.Monster{}

	data := monster.Store()
	monster.WriteFile(data)
	monster.ReadFile()
}
