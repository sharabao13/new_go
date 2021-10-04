package main

import "fmt"

func main() {
	/*
		闭包
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)
			并且该外层函数的返回值就是这个内层函数
			这个内层函数和外层函数的局部变量，统称为闭包结构

			局部变量的生命周期会发生改变，正常的局部变量随着函数的创建而创建，随着函数的结束而结束
			但是闭包结构中外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还是要继续使用

	*/
	//
	i1 := increment() //increment 返回值是一个函数
	res1 := i1()      //将函数调用
	fmt.Println(res1)
}

func increment() func() int { // 外层函数
	i := 0
	fun := func() int { //内层函数
		i++
		return i
	}
	return fun
}