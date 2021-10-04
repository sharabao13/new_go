package main

import "fmt"

func main() {
	/*
			defer 词义 延迟，推迟
			在go语言中，使用defer原件自来延迟一个函数或者方法的执行
			1. defer函数或方法: 一个函数或方法的执行被延迟了
			2. defer的用法
				1. 对象.close(),临时文件被关闭 临时文件的删除
					文件.open()
					defer close()
				2. go语言中关于异常的处理，使用panic和recover（）
					panic函数用于引发恐慌，导致程序的中断执行
					recover函数用于恢复程序的执行，recover（）语法上要求必须在defer中执行
				3.如果多个defer函数 last in first out 先延迟的后执行，后延迟的先执行
		   		4. defer函数传递参数的时候
					defer调用时，就已经传递了参数数据，暂时不执行函数中的代码
				5. defer函数注意点
					1. 当函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
					2. 当执行外围函数中的return语句时，只有其中所有的延迟函数执行完毕后，外围函数才会真正返回
					4. 当外围函数中的代码引发恐慌是，只有其中所有延迟的函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数
	*/
	defer fun1("hello") //也被延迟
	fmt.Println("123")
	defer fun1("world") //延迟
	fmt.Println("456")
}

func fun1(s string) {
	fmt.Println(s)
}
