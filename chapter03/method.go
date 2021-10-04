package main

import "fmt"

func main() {
	/*
		go语言中同时用函数和方法，一个方法就是一个包含了接收者的函数，接受者可以是命名类型或者结构体类型的一个值或者
		是一个指针。所有给定类型的方法属于该类型的方法集

		方法
		方法只是一个函数，它带有一个特殊的介绍器类型，它是在func关键字和方法名之间编写的，接收器可以是struce类型
		或非struct类型，接收方可以在方法内部使用

		定义方法的语法
		func (t type) methodName(parameter list) (return list) {}
		func (接收者) 方法名(参数列表) (返回值列表) {}

		总结 method 同函数类似，区别需要有接受者

		函数的定义
		func  functionName(parameter list) (return list) {}

		对比函数
			1. 意义
				方法: 某个类别的行为功能，需要指定的接收者调用
				函数: 一段独立功能的代码，可以直接调用

			2. 语法
				方法： 方法名称可以相同，只要接收者不同
				函数: 命名不能冲突
	*/
	//创建一个机构提对象
	w1 := worker{"IU", 18} //值类型传递
	w1.study()             //值类型调用

	w2 := &worker{"HCX", 18} //指针类型传递
	w2.study()               //指针类型调用

	w2.rest()
	w1.rest()

	c1 := &cat{"M", 2}
	fmt.Println(*c1)
	c1.printInfo()
	c1.name = "R"
	fmt.Println(*c1)
	c1.printInfo()
	w1.printInfo()
	w2.printInfo()

}

//定义一个结构体
type worker struct {
	name string
	age  int
}

type cat struct {
	name string
	age  int
}

//定义一个方法
func (w worker) study() {
	fmt.Println(w.name, "在学习")
}

func (w *worker) rest() {
	fmt.Println(w.name, "在休息")
}

func (p *cat) printInfo() {
	fmt.Printf("猫的姓名: %s,猫的年龄: %d\n", p.name, p.age)
}

func (p *worker) printInfo() {
	fmt.Printf("工人的姓名: %s,工人的年龄: %d\n", p.name, p.age)
}
