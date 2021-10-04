package main

import "fmt"

func main() {
	/*
		切片slice
			1. 每个切片引用了一个底层数组
			2. 切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
			3. 当想切片中加数据时，如何没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
			4. 切片一旦扩容，就是重新指向一个新的底层数组

	*/
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 6, 7, 8)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 11, 12, 13, 14, 15)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

}
