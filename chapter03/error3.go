package main

import "fmt"

func main() {
	area, err := rectArea(-9.0, -6.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(area)
}

type areaError1 struct {
	msg           string //错误描述
	lenght, width float64
}

func (e *areaError1) Error() string {
	return e.msg
}

func (e *areaError1) lenghtNegative() bool {
	return e.lenght < 0
}

func (e *areaError1) widthNegative() bool {
	return e.width < 0
}

func rectArea(lenght, width float64) (float64, error) {
	msg := ""
	if lenght < 0 {
		msg = "长度小于0"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度小于0"
		} else {
			msg += "宽度也小于0"
		}
	}
	if msg != "" {
		return 0, &areaError1{msg, lenght, width}
	}
	return lenght * width, nil
}
