package main

import "fmt"

func main() {
	/*
		接口
		面向对象世界中的接口一般定义是"接口定义对象的行为",它标识对象该做什么，实现这种行为的方法是针对对象的
		在go中，接口是一组方法签名，当类型为接口中的所有方法提供定义时，它被称为实现接口
		接口指定了类型应该具有的方法，类型决定了如何实现这些方法

		go语言中，接口和类型的实现关系，是非侵入式

		1. 当需要接口类型实现对象时，可以使用任意实现类对象代替
		2. 接口对象不能访问实现类中的属性

		多态: 一个事物的多种形态
			go语言通过接口来模拟多态

			就一个接口的实现
				1. 看成实现本身的类型，能过访问实现类中的属性方法
				2. 看成是对应的接口类型，那只能访问接口中定义的方法

		接口的用法
			1. 一个函数如果接收接口类型作为参数，那么实际上可以传入该接口类型的任意实现类类型对象做为参数
			2. 定义一个类型为实现接口类型，实际上可以赋值为任意实现类的对象
	*/
	// 创建mouse对象
	m1 := mouse{"罗技"}
	fmt.Println(m1.name)

	f1 := flashdisk{"闪迪"}
	fmt.Println(f1.name)

	//调用测试方法
	testInterface(m1) //m1是mouse的结构体类型,mouse已经实现了这个接口的方法，所以m1也是这个接口的实现
	//usb.start(m1)  单独调用某个接口的方法

	testInterface(f1)
	//usb.end(f1)

	var u usb
	u = f1
	u.start()
	//fmt.Println(u.name)  不能访问实现类的字段属性
	f1.deleteData()
	//u.deleteData() 不是所有的usb都是U盘所以U盘新增的实现方法不能调用

}

//1 定义接口
type usb interface {
	start()
	end()
}

//2 实现类
type mouse struct {
	name string
}

type flashdisk struct {
	name string
}

//3. 实现类的方法
func (m mouse) start() {
	fmt.Println(m.name, "鼠标开始工作了...")
}

func (m mouse) end() {
	fmt.Println(m.name, "鼠标结束工作...")
}

func (f flashdisk) start() {
	fmt.Println("U盘开始工作了...")
}
func (f flashdisk) end() {
	fmt.Println("U盘结束工作...")
}
func (f flashdisk) deleteData() {
	fmt.Println(f.name, "U盘删除数据")
}

//4. 测试方法
func testInterface(u usb) { //u = m1
	u.start()
	u.end()
}
