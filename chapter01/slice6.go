package main

import "fmt"

func main() {
	s1 := make([]int, 3, 3)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

}
