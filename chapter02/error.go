package main

import (
	"fmt"
)

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//
	//fmt.Println("main...start...")
	////panic("error")
	//fmt.Println("main...over...")
	err := testEroor()
	fmt.Println(err)
}
func testEroor() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("error")
	return
}
