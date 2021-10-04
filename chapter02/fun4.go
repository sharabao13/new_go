package main

import "fmt"

func main() {
	/*
		函数的返回值  一个函数被调用后，返回给调用处的执行结果，
		调用处需要使用变量接收该结果

		return 语句
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处

		函数多返回值
	*/
	// 1-100 的和
	res := getsum(100)
	fmt.Println("1-100的和", res)
	fmt.Println("---------------------")

	res1, res2 := rec(5, 3)
	fmt.Println("周长是： ", res1, "面积是: ", res2)
	fmt.Println("---------------------")

	res3, res4 := rec(5, 3)
	fmt.Println("周长是： ", res3, "面积是: ", res4)
	fmt.Println("---------------------")
}
func getsum(a int) int { //返回值
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	return sum //必须使用return 语句返回给调用处
}

//函数多返回值,求周长和面积
func rec(a, b float64) (float64, float64) {
	peri := (a + b) * 2
	area := a * b
	return peri, area
}

func rec1(a, b float64) (peri, area float64) {
	peri = (a + b) * 2
	area = a * b
	return
}
