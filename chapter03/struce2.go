package main

import (
	"fmt"
)

func main() {
	/*
		匿名结构体和匿名字段
		匿名结构体： 没有名字的结构体
			在创建匿名结构体时，同时创建对象
		匿名字段： 一个结构体的字段没有名字
	*/
	p1 := students{"HCX", 18}
	fmt.Println(p1.name, p1.age)

	//匿名函数
	func() {
		fmt.Println("hello world...")
	}()

	//匿名结构体
	p2 := struct {
		name string
		age  int
	}{
		name: "BAO",
		age:  18,
	}
	fmt.Println(p2.name, p2.age)

	//调用结构体匿名字段
	p3 := worker{
		"IU",
		18,
	}
	fmt.Println(p3.string, p3.int)
}

//正常结构体
type students struct {
	name string
	age  int
}

type worker struct {
	string //匿名字段
	int    //匿名字段，默认使用数据类型做为名字调用，那么匿名字段的类型不能重复，否则会冲突
	//string
}
