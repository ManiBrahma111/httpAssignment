package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	os.OpenFile("sample.txt", 1, 0)
	f.Write([]byte("hello world"))

	str := []byte("new string is being added")
	os.WriteFile("sample.txt", str, 14609)
	info, _ := f.Stat()
	fmt.Println(info)
	str = []byte("Hello world")
	os.WriteFile("sample.txt", str, 14609)

	st, _ := os.ReadFile("sample.txt")
	fmt.Println(string(st))
	defer f.Close()
}
