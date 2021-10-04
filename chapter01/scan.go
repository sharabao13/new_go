package main

import (
	"fmt"
)

func main() {
	// 从键盘输入

	var (
		name string
		age  int
	)
	fmt.Println("input you name and age")
	//fmt.Println("input you age")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)

	var (
		n1 string
		a1 int
	)
	fmt.Println("请输入姓名: ")
	fmt.Scan(&n1)
	fmt.Println("请输入年龄: ")
	fmt.Scan(&a1)
	fmt.Printf("您的姓名是: %s 您的年龄是: %d\n", n1, a1)
}
