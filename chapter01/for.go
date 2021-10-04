package main

import "fmt"

func main() {
	/*
			for 循环语法
		    标准写法
		    for 表达式1;表达式2;表达式3; {
		     	循环体
		     }
		     执行顺利  表达式1 执行1次 通知用来做变量初始化
		              表达式2 bool循环的条件
		              表达式3 变量的编号
		              循环体 执行代码

		    其他写法
		    省略表达式1 表达3
		     for 表达式2{
		    	循环体
		        变量控制
		     }

		    同时省略表达式1 表达式2 表达式3
		    相当于作用在while(true)上面
		    条件永远为真
	*/
	//标准写法
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello World")
	}
	//省略表达式1 3 写法
	i1 := 1
	for i1 <= 5 {
		fmt.Println("hcx")
		i1++
	}
	/*
			for练习题
		    1. 打印58 -23 的数字
		    2. 求1 - 100 的和
		    3. 打印1 -100 内 能被3整除的，但不能被5整除的数字，统计被打印数字的个数，每行打印5个
	*/
	//练习1
	for num1 := 58; num1 >= 23; num1-- {
		fmt.Println(num1)
	}

	//练习2
	sum := 0
	for i2 := 1; i2 <= 100; i2++ {
		sum += i2
	}
	fmt.Println(sum)

	//练习3
	total := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i)
			fmt.Printf("\t")
			total++
			if total%5 == 0 {
				//fmt.Println(i)
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println(total)
}
