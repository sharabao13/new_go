package main

import (
	"fmt"
)

func main() {
	/*
		空接口 interface{}
			不包含任何的方法，正因为如此，所有类型都实现了空接口，因此空接口可以存储任意类型的数值
	*/
	b1 := b{name: "xxx"}
	c1 := c{age: 18}
	//空接口可以赋值给任意类型的数据
	var a1 a = "空接口"
	var a2 a = 18
	var a3 a = b1
	var a4 a = c1
	fmt.Println(a1, a2, a3, a4)
	//空接口函数调用
	test(a1)
	test(a2)
	test(a3)
	test(18)
	test3("ruby")
	m1 := make(map[string]interface{})
	m1["name"] = "arx"
	m1["age"] = 18
	m1["friend"] = b{name: "hcx"}
	fmt.Println(m1)

	s1 := make([]interface{}, 0, 10)
	s1 = append(s1, a1, a2, a3, a4)
	fmt.Println(s1)

	test4(s1)

}

type a interface {
}

type b struct {
	name string
}
type c struct {
	age int
}

func test(a0 a) {
	fmt.Println(a0)
}
func test3(a0 interface{}) {
	fmt.Println(a0)
}

func test4(slice []interface{}) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
