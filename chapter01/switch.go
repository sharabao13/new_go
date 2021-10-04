package main

import (
	"fmt"
)

func main() {
	// switch 语句
	//switch var1 {
	//case val1:
	//	...
	//case val2:
	//	...
	//default:
	//	...
	//}
	/*
			注意事项:
		       1. switch 可以作用在其他类型的，case后的数值必须和switch的变量类型一致
		       2. case是无序的
		       3. case后的数值是唯一的
		       4. default语句是可选的
		       5. 省累switch后的变量，相当于直接作用在true上
		       6. switch 可以初始化语句
	*/
	finger := 7
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Invalid")
	}

	//表达式case
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// 没有表达式的switch
	num := 75
	switch {
	case num >= 0 && num < 60:
		fmt.Println("number is greater 0 and less than 60")
	case num >= 60 && num <= 100:
		fmt.Println("number is grate 60 and less than 100")
	case num > 100:
		fmt.Println("number is greate 100")
	}

	switch num1 := 75; {
	case num1 < 50:
		fmt.Println("number1 is less than 50")
		fallthrough
	case num1 < 100:
		fmt.Println("number1 is less than 100")
		fallthrough
	case num1 < 60:
		fmt.Println("number1 is less than 60")
		fallthrough
	case num1 < 200:
		fmt.Println("number 1 less than 200")
	}

	//var score int
	//fmt.Println("请输入成绩")
	//fmt.Scan(&score)
	//switch  {
	//case score >= 90:
	//	fmt.Println("A")
	//case score >= 80:
	//	fmt.Println("B")
	//case score >= 70:
	//	fmt.Println("c")
	//case score >= 60:
	//	fmt.Println("D")
	//default:
	//	fmt.Println("E")
	//}

	var (
		num2 int
		num3 int
		oper string
	)
	fmt.Println("请输入一个整数: ")
	fmt.Scanln(&num2)
	fmt.Println("请输入另一个整数: ")
	fmt.Scanln(&num3)
	fmt.Println("请输入一个操作数: ")
	fmt.Scanln(&oper)
	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d\n", num2, num3, num2+num3)
	case "-":
		fmt.Printf("%d - %d = %d\n", num2, num3, num2-num3)
	case "*":
		fmt.Printf("%d * %d = %d\n", num2, num3, num2*num3)
	case "/":
		fmt.Printf("%d / %d = %d\n", num2, num3, num2/num3)
	}
	fmt.Printf("main... over...")
}
