package main

import "fmt"

func main() {
	/*
		递归函数 一个函数自己调用自己
			递归函数要有一个出口，逐渐的像出口靠近
	*/
	//求1-5的和
	sum := getsum(5)
	fmt.Println("1-5的和", sum)

	//求5阶乘
	fac := factorial(10)
	fmt.Println("5的阶乘", fac)

	//fibonacci
	fibo := fibonacci(12)
	fmt.Println("fibo: ", fibo)
}
func getsum(n int) int {
	if n == 1 {
		return 1
	}
	return getsum(n-1) + n
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
