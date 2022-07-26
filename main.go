package main

import (
	"fmt"
)

type j interface {
	walk()
}

type i interface {
	j
	greet()
	add(a, b int)
}

type s struct {
	name string
	age  int
}

func (str s) greet() {
	fmt.Println("Hi how are you")
}
func (str s) walk() {
	fmt.Println("Hi i am walking")
}
func (str s) add(a, b int) {
	fmt.Println(a + b)
}

//goland:noinspection GoUnusedCallResult
func main() {
	fmt.Println("Hello world")
	fmt.Println("I am building a calci\n enter two numbers")
	//var a, b int
	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	//fmt.Println("sub :", sub(a, b))
	//fmt.Println("add :", add(a, b))
	//fmt.Println("multi :", multi(a, b))
	//fmt.Println("div :", div(a, b))
	//fmt.Println("mod :", projext.Modu(a, b))
	var t = s{
		name: "Mani",
		age:  16,
	}
	t.walk()
	t.greet()
	t.add(5, 10)
}
