package main

import (
	"fmt"
)

func main() {
	/*
		匿名函数 没有名字的函数
		定义一个匿名函数，直接进行调用，通常只能使用一次在函数末尾直接在()
		也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次

		注意事项
			1. go语言支持函数式编程
			2. 将匿名函数作为另一个函数的参数，回调函数
			3. 将匿名函数作为另一个函数的返回值，可以形成闭包
	*/
	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数 ")
	}()

	//赋值给变量使用
	fun1 := func() {
		fmt.Println("我是一个匿名函数 ")
	}
	fun1()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res1)
}
