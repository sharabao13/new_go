package main

import "fmt"

func main() {
	/*
			布尔类型 true false
		    零值默认为false
	*/
	var emptyBool bool
	isTrue := false
	isFalse := true
	fmt.Println(emptyBool, isFalse, isTrue)
	fmt.Println("-----------------------")

	//操作
	//逻辑运算 与 &&或 ||非 ！
	// 与运算  //都为真才真 一个假都为假
	fmt.Println(true && true)   //真真   为真
	fmt.Println(true && false)  //真假  为假
	fmt.Println(false && true)  //假真  为假
	fmt.Println(false && false) //假加  为假
	fmt.Println("-----------------------")

	// 或运算  一个真都为真 都为假才为假
	fmt.Println(true || true)   //真真   为真
	fmt.Println(true || false)  //真假  为真
	fmt.Println(false || true)  //假真  为真
	fmt.Println(false || false) //假加  为假
	fmt.Println("-----------------------")

	//！ 非运算  取反
	fmt.Println(!true)
	fmt.Println(!false)
	//关系运算 等 == 不登  =
	fmt.Println("-----------------------")

	//关系运算
	//相等
	fmt.Println(isTrue == isFalse)
	fmt.Println(isTrue == emptyBool)
	fmt.Println(isFalse == emptyBool)
	//不相等
	fmt.Println(isTrue != isFalse)
	fmt.Println(isTrue != emptyBool)
	fmt.Println(isFalse != emptyBool)

	//占位符 %t
	fmt.Printf("isFalse=%T,isFalse=%t\n", isFalse, isFalse)

}
