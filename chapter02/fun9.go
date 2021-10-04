package main

import "fmt"

func main() {
	/*
		回调函数
		高阶函数 根据go语言的数据类型特点，可以将一个函数作为另一个函数的参数

		fun1() fun2()
		将fun1()函数作为fun2这个函数的参数
			fun2函数  高阶函数
				接收了一个函数作为参数的函数，高阶函数
			fun1函数  回调函数
				作为另一个函数的参数的函数，回调函数
	*/
	//加法操作
	add1 := oper(1, 2, add) //add函数作为参数传给oper函数 在28行调用
	fmt.Println(add1)

	//减法操作
	sub1 := oper(10, 2, sub) //add函数作为参数传给oper函数 在28行调用
	fmt.Println(sub1)

	//匿名函数作为变量 乘法
	rest1 := func(a, b int) int {
		return a * b
	}
	multi := oper(2, 9, rest1)
	fmt.Println(multi)

	//匿名函数直接传递给函数 //除法
	div1 := oper(9, 2, func(a, b int) int {
		if b == 0 {
			fmt.Println("被除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(div1)
}

// 加法函数
func add(a, b int) int {
	return a + b
}

//减法操作
func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(a, b int) int) int {
	res1 := fun(a, b) //add函数回调函数
	return res1
}
