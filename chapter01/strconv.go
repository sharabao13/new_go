package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包 字符串和基本类型之前的转换
		string strconv
		parse 字符串转换为其他数据类型
		format 将其他数据类型转换为字符串
	*/
	//字符串转布尔类型
	s1 := "true"
	t1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\t%t\n", t1, t1)
	//bool 转 string
	s2 := strconv.FormatBool(t1)
	fmt.Printf("%T\t%s\n", s2, s2)

	//str -》 int整数
	s3 := "100" // base 转换整形的进制8 10 2 bitsize 最大位数
	i1, err := strconv.ParseInt(s3, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\t%d\n", i1, i1)

	// int -=> string base 进制
	s4 := strconv.FormatInt(i1, 10)
	fmt.Printf("%T\t%s\n", s4, s4)

	//整形转换常用atoi itoa
	i2 := 20
	s5 := strconv.Itoa(i2)
	fmt.Printf("%T\t%s\n", s5, s5)

	s6 := "18"
	i3, err := strconv.Atoi(s6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\t%d\n", i3, i3)
}
