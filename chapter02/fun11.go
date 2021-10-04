package main

import "fmt"

func main() {
	fmt.Println(calc("add", 1, 2, 4, 5))

	answer := Recursive(4)
	fmt.Printf("Recursive: %d\n", answer)

	tower("A", "B", "C", 3)

	a, b, c := "A", "B", "C"
	var val int
	fmt.Println("请输入要移动的盘子数（1-20之间）：N=")
	fmt.Scanf("%d", &val)
	hannuota(val, a, b, c)

}
func hannuota(n int, A, B, C string) {
	if n < 1 || n > 20 {
		return
	}
	if n == 1 {
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
	} else {
		hannuota(n-1, A, C, B)
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
		hannuota(n-1, B, A, C)
	}
}

func tower(a, b, c string, layer int) {
	fmt.Println("*****************")
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	// a n-1 借助 c 移动到 b
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	// b n-1 借助 a 移动到 c
	tower(b, a, c, layer-1)
}

func calc(oper string, a, b int, args ...int) int {
	switch oper {
	case "add":
		sum := 0
		sum = a + b
		for _, v := range args {
			sum += v
		}
		return sum
	}
	return 0
}
func Recursive(number int) int {
	if number == 1 {
		return number
	}
	return number + Recursive(number-1)
}
