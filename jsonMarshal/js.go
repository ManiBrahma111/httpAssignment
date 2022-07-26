package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Num    int    `json:"num"`
	Branch string `json:"branch"`
}

func main() {
	var s = Student{
		Name:   "mani",
		Num:    100,
		Branch: "ECE",
	}
	op, err := json.Marshal(s)
	//to apply indentation we use MarshalIndent
	opp, errr := json.MarshalIndent(s, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}
	// return type is byteArray should convert it into string
	fmt.Println(string(op))
	if errr != nil {
		fmt.Println(errr)
	}
	fmt.Println(string(opp))
	val := `{"Name": "sam","Num": 125,"Branch": "CSE"}`
	var mapp Student
	er := json.Unmarshal([]byte(val), &mapp)

	if er != nil {
		fmt.Println(err)
	}
	fmt.Println(mapp)

}
