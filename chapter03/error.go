package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		error： 内置的数据类型，内置的接口
			定义方法Error() string

		使用go语言提供好的包
			errors包下的函数new 创建一个error对象
			fmt包下的errorf（）函数
	*/

	err1 := errors.New("错误类型")
	fmt.Println(err1)

	err2 := fmt.Errorf("错误的信息码: %d\n", 100)
	fmt.Println(err2)

	f, err := os.Open("D:\\go\\src\\github.com\\sharabao13\\new_go\\chapter03\\test.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err) //open test.txt: no such file or directory
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")

	err3 := checkAge(-10)
	if err3 != nil {
		fmt.Println(err3)
	}
}
func checkAge(age int) error {
	if age < 0 {
		err := errors.New("年龄不合法")
		return err
	}
	fmt.Println("年龄为: ", age)
	return nil
}
