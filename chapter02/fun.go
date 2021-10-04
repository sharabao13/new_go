package main

import "fmt"

func main() { //程序的入口，一个特殊的函数
	/*
		函数 function
			一、概念
				具有特定功能的代码，可以被多次调用执行
			二、意义
				1.可以避免重复的代码
				2.增强程序的扩展性
			三、使用
				1. 函数的定义，函数的声明
				2. 函数的调用，就是执行函数中的代码
			四、语法
				func funcName (parametername type1,parametername type2) (output1 type1,output2 type2) {
					//这里是处理逻辑的代码
					//返回值
					return value1,value2
				}
			五、函数的参数
				参数的使用
					形式参数： 定义函数时，用于接收外部传入的数据，
					实际参数： 调用函数时，传给行参的实际的数据，叫做实际参数
				函数调用
					1. 函数名称必须匹配
					2. 实参与形参必须一一对应，顺序，个数，类型
	*/
	// 求1-100 的和
	getSum(100)

	// 求1-200 的和
	getSum(200)
}

// 没有参数
//func getSum() {
//	sum := 0
//	for i := 1; i <= 100; i++ {
//		sum += i
//	}
//	fmt.Println("1-100的和", sum)
//}
//参数
func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和: %d\n", n, sum)
}
