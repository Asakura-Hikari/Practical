package main

import "fmt"

type Dog struct {
	name   string
	age    int
	weight float64
}

func (d Dog) say() string {
	str := fmt.Sprint("小狗的名字是:", d.name, "\n年龄是:", d.age, "\n体重是: ", d.weight)
	return str
}

func main() {
	var dog Dog = Dog{"白藏主",2000,50.00}
	fmt.Println(dog.say())
}