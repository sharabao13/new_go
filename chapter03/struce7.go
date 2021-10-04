package main

import "fmt"

func main() {
	t1 := test1{"aaa"}
	fmt.Println(t1.name)
	t2 := t1
	t2.name = "bbb"
	fmt.Println(t1.name)
	fmt.Println(t2.name)

	var t3 *test2
	t4 := test2{"CCCC"}
	t3 = &t4
	fmt.Println(t4.name)
	fmt.Println(t3.name)
	fmt.Println("-------------")
	t3.name = "dddddddddddddd"
	fmt.Println(t4.name)
	fmt.Println(t3.name)
}

type test1 struct {
	name string
}

type test2 struct {
	name string
}
