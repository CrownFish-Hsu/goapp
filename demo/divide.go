package demo

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("param error")
	}

	return a / b, nil
}

func TestDivide() {
	num1 := 10.0
	num2 := 2.0
	// 调用 divide 函数进行除法运算
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f 除以 %f 的结果是: %f\n", num1, num2, result)
	}

	num3 := 5.0
	num4 := 0.0
	result, err = divide(num3, num4)
	if err != nil {
		fmt.Printf("%f 除以 %f 的结果是: %s\n", num3, num4, err.Error())
	} else {
		fmt.Printf("%f 除以 %f 的结果是: %f\n", num3, num4, result)
	}
}
