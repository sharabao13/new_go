package main

import "fmt"

func main() {
	/*
		变量定义 var

	*/
	//定义单个变量
	var fristVar string
	fristVar = "one"
	fmt.Println(fristVar)

	//定义多个同类型的变量
	var name, user string
	name = "hcx"
	user = "gfb"
	fmt.Println(name, user)

	//定义多个同类型变量并赋值
	var userName, userPass string = "hcx", "gfb"
	fmt.Println(userName, userPass)

	//定义多个不同类型的变量
	var (
		userPhone int    = 31
		gender    string = "std"
	)
	fmt.Println(userPhone, gender)

	var (
		s1 string
		i1 int
	)
	s1 = "hcx"
	i1 = 18
	fmt.Println(s1, i1)

	//定义变量 赋值类型自动推导
	var (
		s2 = "hcx"
		i2 = 18
	)
	fmt.Println(s2, i2)

	var s3, i3 = "hcx", 18
	fmt.Println(s3, i3)

	//变量简短声明
	isLove := false
	fmt.Println(isLove)
}
