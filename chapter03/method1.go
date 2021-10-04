package main

import "fmt"

func main() {
	/*
		方法模拟继承性
	*/
	//创建父类对象
	p1 := person{"HCX", 18}
	//调用父类的方法
	p1.eat()

	//创建子类对象
	s1 := student{person{"GFB", 18}, "厦门大学"}
	//子类调用父类的方法
	s1.eat()
	//子类访问子类新增的方法
	s1.study()
	//子类访问自己重新的方法
	s1.eat()
}

// 定义父类结构体
type person struct {
	name string
	age  int
}

// 定义子类结构体
type student struct {
	person //匿名字段
	school string
}

//定义父类的方法
func (p person) eat() {
	fmt.Println(p.name, "父类的方法在吃饭")
}

//定义子类新增的方法
func (s student) study() {
	fmt.Println("在学习")
}

//定义子类重新父类的方法
func (p student) eat() {
	fmt.Println(p.name, "子类的的方法在吃汉堡")
}
