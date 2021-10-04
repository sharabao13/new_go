package main

import "fmt"

func main() {
	/*
			指针	存储另一个变量的内存地址的变量
			一个指针可以指向任何一个值的内存地址它指向那个值的内存地址
		 	获取指针的地址符是 & 返回对应变量的内存地址

			声明一个指针的语法
			var 变量名 *type
	*/

	a := 10
	fmt.Printf("a的值是%d\n", a)
	fmt.Printf("a的类型是%T\n", a)
	fmt.Printf("a的内存地址%p\n", &a)
	fmt.Println("----------------------")

	// 创建一个指针变量，用于存储变量a的地址
	var p1 *int
	p1 = &a
	fmt.Printf("p1的值是%v\n", p1)
	fmt.Printf("p1的类型是%T\n", p1)
	fmt.Printf("p1的内存地址%p\n", &p1)
	fmt.Printf("p1的内存地址对应的数组%d\n", *p1)

	// 操作变量，更改数值，并不会改变地址
	a = 100
	fmt.Println(a)
	fmt.Printf("a的内存地址%p\n", &a)

	//通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)
	fmt.Printf("a的内存地址%p\n", &a)

	//指针的指针
	var p2 **int
	fmt.Println(p2) //空指针
	//赋值
	p2 = &p1
	fmt.Println(p2)
	fmt.Println("p2对应内存的内存地址是", **p2)
}
