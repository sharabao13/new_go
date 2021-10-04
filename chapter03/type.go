package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		type 关键字
			type 是go语语法里重要且常用的关键字，type绝不只是对应于c/c++中的typedef，
			1. 结构体 type struceName struce {}
			2. 接口 type interfaceName interface {}
			3. 定义其他新类型  type 类型名 Type


	*/

	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)
	//i1 = i2
	//fmt.Println(i1,i2) cannot use i2 (type int) as type myint in assignment

	var s1 mystring
	var s2 = "world"
	s1 = "hello"
	fmt.Println(s1, s2)

	rest := fun1()
	fmt.Println(rest(10, 20))
}

//定义新的 类型
type myint int
type mystring string

//类型别名
type myint2 = int //不是重新定义新的数据类型，只是给int起别名,和int通用

//定义函数
type myfunc func(a, b int) string

func fun1() myfunc {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}
