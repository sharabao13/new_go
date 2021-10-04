package main

import "fmt"

func main() {
	//switch  break fallthrough
	/*
			break 可以使用在switch中，也可使用在for循环中
		          switch中 强制介绍case，从而结束switch
		    fallthrough 贯通后续case fallthrough 只能在语句的最好一行
	*/

	n := 2
	switch n {
	case 1:
		fmt.Println("i am hxc")
		fmt.Println("i am hxc")
		fmt.Println("i am hxc")
	case 2:
		fmt.Println("i am bao")
		fmt.Println("i am bao")
		break
		fmt.Println("i am bao")
	case 3:
		fmt.Println("i am iu")
		fmt.Println("i am iu")
		fmt.Println("i am iu")
	}
	fmt.Println("main...over...")

	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
		fallthrough
	case 4:
		fmt.Println("第四季度")
	}
}
