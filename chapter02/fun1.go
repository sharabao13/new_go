package main

import "fmt"

func main() {
	/*
			函数的多个参数
		    函数调用
				1.函数名“ 声明的函数名和调用的函数名要统一
				2.实参必须严格匹配形参： 顺序，个数，类型
	*/
	getAdd(10, 20)
	getAdd1(1, 2)
	//fun1(2, 3, 10)  //实参类型不匹配
	//fun1(1.22,2,3333,"str")  //参数不匹配
	fun1(3.13312, 2.777, "hcx")
}
func getAdd(a int, b int) {
	sum := a + b
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}

func getAdd1(a, b int) {
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
}

func fun1(a, b float64, c string) {
	fmt.Printf("%.2f,%.2f,%s\n", a, b, c)
}
