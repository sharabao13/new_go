package main

import "fmt"

func main() {
	/*
		go语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型
		结构体 由一系列具有相同类型或者不同类型的数据构成的数据集合
		一旦定义了结构体类型，它就能用于变量的声明
		结构体格式
		type struce_variable_type struce {
			member definition;
			member definition;
		}

	*/
	// 初始化结构体
	//方法1
	var P1 Persion
	fmt.Println(P1)
	P1.name = "HCX"
	P1.age = 18
	P1.gender = "女"
	P1.address = "泉州市"
	fmt.Printf("%s,%d,%s,%s\n", P1.name, P1.age, P1.gender, P1.address)

	// 方法二
	p2 := Persion{}
	p2.name = "GFB"
	p2.age = 18
	p2.gender = "男"
	p2.address = "泉州市"
	fmt.Printf("%s,%d,%s,%s\n", p2.name, p2.age, p2.gender, p2.address)

	//方法三
	p3 := Persion{name: "TAE", age: 18, gender: "女", address: "厦门"}
	fmt.Println(p3)
	p4 := Persion{
		name:    "IU",
		age:     18,
		gender:  "女",
		address: "厦门",
	}
	fmt.Println(p4)

	p5 := Persion{"Lisa", 18, "女", "厦门"}
	fmt.Println(p5)
}

//定义结构体
type Persion struct {
	name    string
	age     int
	gender  string
	address string
}
