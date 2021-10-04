package main

import "fmt"

func main() {
	/*
		接口模拟多态
	*/
	//step1 创建父类对象
	a1 := animal{"ma", 18}

	//step2 创建子类对象cat
	c1 := cats{animal{name: "mao", age: 18}, "白色"}

	//step3 创建子类对象dog
	d1 := dogs{animal{"gou", 18}}

	//step4 模拟多态
	//父类对象的实现
	a1.eat()
	a1.sleep()
	//a1.lookDoor() 不能使用子类对象的行为方法

	//子类对象cat的实现
	c1.eat()
	c1.sleep()

	//子类对象dog的实现
	d1.eat()
	d1.sleep()
	d1.lookDoor()

	//函数调用
	testAnimal(a1)
	testAnimal(c1)
	testAnimal(d1)
}

//定义一个父类的实现类接口
type animals interface {
	eat()
	sleep()
}

//定义父类对象 animal
type animal struct {
	name string
	age  int
}

//定义子类对象 cat
type cats struct {
	animal
	color string
}

//定义子类对象 dog
type dogs struct {
	animal
}

//定义父类animal 的行为方法eat
func (a animal) eat() {
	fmt.Println(a.name, "吃饭")
}

//定义父类animal 的行为方法sleep
func (a animal) sleep() {
	fmt.Println(a.name, "在睡觉")
}

//定义子类dogs的行为fangfa  lookDoor
func (d dogs) lookDoor() {
	fmt.Println(d.name, "在看家")
}

//创建测试函数
func testAnimal(a animals) {
	a.sleep()
	a.eat()
}
