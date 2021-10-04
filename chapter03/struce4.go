package main

import "fmt"

func main() {
	/*
		结构体比较
		结构体是值类型，如果每个字段具有可比性，则是可比较的。如果它们对应的字段相等，则认为两个结构体变量是相等的。
	*/
	name1 := name{"leborn", "james"}
	name2 := name{"leborn", "james"}
	if name1 == name2 {
		fmt.Println("name1 is equal name2")
	} else {
		fmt.Println("name1 is not equal name2")
	}
	name3 := name{"GUAN", "FUBAO"}
	name4 := name{"leborn", "JAMES"}

	if name3 == name4 {
		fmt.Println("name3 is equal name4")
	} else {
		fmt.Println("name3 is not equal name4")
	}
}

type name struct {
	firstName string
	lastName  string
}
