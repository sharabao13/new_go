package main

import "fmt"

func main() {
	/*
		结构体嵌套 一个结构体中的字段，是另一个结构体的类型
	*/
	//创建book的对象
	b1 := book{
		name:  "GO",
		price: 76.33,
	}
	s1 := students{
		name: "HCX",
		age:  18,
		book: b1,
	}
	//结构体嵌套的对象调用方法s1.b1.
	fmt.Printf("姓名: %s,年龄: %d,看的书名: %s,书的价格是%.2f\n", s1.name, s1.age, s1.book.name, s1.book.price)

	s2 := students{
		name: "BAO",
		age:  18,
		book: book{
			name:  "Python",
			price: 89.66,
		},
	}
	fmt.Printf("姓名: %s,年龄: %d,看的书名: %s,书的价格是%.2f\n", s2.name, s2.age, s2.book.name, s2.book.price)

	b2 := book{
		"JAVA",
		66.24,
	}
	s3 := students1{
		name: "IU",
		age:  18,
		book: &b2,
	}
	fmt.Println(b2)
	fmt.Println(s3.book.name, s3.book.price)
	s3.book.name = "C++"
	fmt.Println(b2)
	fmt.Println(s3.book.name, s3.book.price)
}

//书的信息
type book struct {
	name  string
	price float64
}

//学生信息
type students struct {
	name string
	age  int
	book book
}

type students1 struct {
	name string
	age  int
	book *book
}
