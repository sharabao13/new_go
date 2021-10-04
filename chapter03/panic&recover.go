package main

import (
	"fmt"
)

func main() {
	/*
			golang中引入两个内置函数panic和recover来触发和终止异常处理流程，同时引入关键字defer来言延迟执行defer
			后面的函数

			一直等到包含defer语句的函数执行完毕时，延迟函数才会被执行，而不管包含defer语句的函数是通过retrun的正常结束
			。你可以在 函数中执行多条defer语句，他们的执行顺序是与声明顺序相反

		    当程序运行时，如果遇到空指针、下标越界，或显示调用panic函数等情况， 则先触发panic函数的执行，然后调用延迟函数
			。调用者继续传递panic，因此该过程一直在调用栈中重复发生，函数停止执行，调用延迟执行函数，，如果一路
			在延迟函数中没有recover函数的调用，则会达到该协成的起点，改协成结束，然后终止其他所有协成，包含主协成

			panic
				1. 内建函数
				2. 假如函数中书写了panic语句，会终止其后要执行的代码，在painc所在函数内如果存在要执行的defer函数列表，
					按照defer的逆序执行
				3. 返回函数的调用者，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照derfer
					的逆序执行，这里的defer有点类似try-catch-finally 中的finally
				4. 直到goroutine整个退出，并报告错误

		 	recover
				1. 内建函数
				2. 用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
				3. 一般的调用建议
					a. 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码执行
					b. 可以获取通过panic传递的error
	*/
	funA()
	defer myprint("defer main: 3....")
	funB()
	defer myprint("defer main: 4....")
	fmt.Println("main...over...")
}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是一个函数funA")
}

func funB() { //外围函数
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复")
		}
	}()
	fmt.Println("我是一个函数funB")
	defer myprint("defer funB(); 1....")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			panic("funB函数painc")
		}
	}
	defer myprint("defer funB(): 2....")
}
