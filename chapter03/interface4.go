package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		接口断言
		因为空接口interface{}没有定义任何函数，因此go中所有类型都实现了空接口
		当一个函数的形参是interface{}，那么在函数中，需要对形参进行断言，从而得到它的真实类型

		语法格式
		// 安全类型断言
			<目标类型的值>,<布尔参数> := <表达式>.(目标类型)

		//非安全断言
			<目标类型的值> := <表达式>.(目标类型)

		方式一
			1. interface := 接口对象.(实际类型) //不安全 会panic
			2. interface,ok := 接口对象.(实际类型) 安全

		方式二 switch
			switch interface := 接口对象.(type)
			case 实际类型:
				...
			case 实际类型
				...
	*/
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)
	fmt.Println("_____________________")

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())
	fmt.Println("_____________________")
	//接口函数调用
	testShape(t1)
	testShape(s1)
	testShape(s2)
	fmt.Println("_____________________")
	getType(t1)
	getType(c1)
	getType(s1)
	getType(s2)
	fmt.Println("_____________________")

	var t2 *Circle = &Circle{4}
	getType(t2)

	var t3 *Triangle = &Triangle{1, 2, 3}
	getType(t3)
	fmt.Println("_____________________")

	getType1(c1)
	getType1(s1)
	getType1(t1)

}

//定义形状接口
type Shape interface {
	peri() float64
	area() float64
}

//定义三角形实现类
type Triangle struct {
	a, b, c float64
}

//三角形实现方法
func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

//定义圆形的实现类
type Circle struct {
	radius float64
}

//定义圆形的实现方法
func (c Circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 1)
}

//定义一个接口函数
func testShape(s Shape) {
	fmt.Printf("周长是: %.2f,面积是: %.2f\n", s.peri(), s.area())
}

func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是是", ins.radius)
	} else if ins, ok := s.(*Circle); ok {
		fmt.Println("是圆形，半径是是", ins.radius)
	} else {
		fmt.Println("类型错误")
	}
}

func getType1(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("是圆形，半径是是", ins.radius)
	default:
		fmt.Println("type2类型错误")
	}
}
