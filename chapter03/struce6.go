package main

import "fmt"

func main() {
	/*
		面向对象
		go语言的结构体嵌套
			1. 模拟继承性
				type A struce {
					file
				}

				type A struce {
					A //匿名字段
				}

		1. 模拟聚合关系
				type C struce {
					file
				}

				type D struce {
					c C //聚合关系
				}
	*/
	//创建父类对象
	s1 := person1{"HCX", 18}
	fmt.Println(s1.name, s1.age)

	//创建子类对象
	s2 := student1{person1{"IU", 18}, "厦门大学"}
	fmt.Println(s2.person1.name, s2.person1.age, s2.school)

	var s3 student1
	s3.person1.name = "BAO"
	s3.person1.age = 18
	s3.school = "北京大学"
	fmt.Println(s3.person1.name, s3.person1.age, s3.school)

	s3.name = "ruby"
	s3.age = 18
	s3.school = "清华大学"
	fmt.Println(s3.person1.name, s3.person1.age, s3.school)

	/*
		s3.Person.name -------> s3.name
		student11 将person1 作为匿名字段，那么person1中的字段在studen1中将做为提升字段
		提升字段作用 能直接访问

	*/
}

//定义父类结构体
type person1 struct {
	name string
	age  int
}

//定义子类对象
type student1 struct {
	person1
	school string
}
