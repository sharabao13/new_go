package main

import "fmt"

func main() {
	/*
		接口嵌套
	*/
	var c1 cats1 = cats1{}
	c1.tests1()
	c1.tests2()
	c1.tests3()
	fmt.Println("______________")

	var a1 A = c1
	a1.tests1() // 接口A 的实现
	fmt.Println("______________")

	var a2 B = c1
	a2.tests2() // 接口A 的实现
	fmt.Println("______________")

	var a3 C = c1
	a3.tests1()
	a3.tests2() // 接口C 的实现
	a3.tests3()
	fmt.Println("______________")
}

type A interface {
	tests1()
}

type B interface {
	tests2()
}

type C interface {
	A
	B
	tests3()
}

type cats1 struct {
}

func (t1 cats1) tests1() {
	fmt.Println("test1...")
}

func (t1 cats1) tests2() {
	fmt.Println("test2...")
}

func (t1 cats1) tests3() {
	fmt.Println("test2...")
}
