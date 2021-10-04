package main

import "fmt"

func main() {
	////if 语法
	//var evenNum int
	//// 判断奇偶数
	//fmt.Println("请输入一个数字:")
	//fmt.Scan(&evenNum)
	//if evenNum > 0 {
	//	if (evenNum % 2) == 0 {
	//		fmt.Println("输入的数字为偶数")
	//	}  else {
	//		fmt.Println("输入的数字为奇数")
	//	}
	//}else {
	//	fmt.Println("输入的数字需大于0")
	//}

	if num := 11; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	num1 := 10
	if num1 <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num1 > -51 && num1 < 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

	studentScore := 10
	if studentScore > 60 {
		if studentScore >= 60 && studentScore < 70 {
			fmt.Println("student score is C")
		} else if studentScore >= 70 && studentScore < 80 {
			fmt.Println("student score is B")
		} else if studentScore >= 80 && studentScore < 90 {
			fmt.Println("student score is A")
		} else if studentScore >= 90 && studentScore <= 100 {
			fmt.Println("student score is A+")
		} else {
			fmt.Println("Invalid score")
		}
	} else {
		fmt.Println("student core is D")
	}
}
