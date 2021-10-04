package main

import "fmt"

func main() {
	p1 := persion{
		id:   1,
		name: "BAO",
		age:  18,
		addr: "XM",
	}
	fmt.Println(p1)
	fmt.Printf("%p\t%T\n", &p1, p1)

	//值传递
	p2 := persion{
		id:   2,
		name: "HCX",
		age:  18,
		addr: "XM",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%p\t%T\n", &p2, p2)

	//结构体指针
	var pp1 *persion
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p\t%T\n", pp1, pp1)
	fmt.Println(*pp1)

	//使用new()函数，go语言中专门用于创建某种类型的指针
	// 通过new创建的指针不是nil
	pp2 := new(persion)
	fmt.Println(pp2)
	//(*pp2).name
	pp2.name = "TAE"
	pp2.age = 18
	pp2.id = 1
	pp2.addr = "XM"
	fmt.Println(*pp2)
}

type persion struct {
	id   int
	name string
	age  int
	addr string
}
