package demo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TestCalculator() {
	// 创建一个读取器，用于从标准输入读取用户输入
	reader := bufio.NewReader(os.Stdin)

	// 提示用户输入第一个数字
	fmt.Print("请输入第一个数字: ")
	num1Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出现错误:", err)
		return
	}
	num1Str = strings.TrimSpace(num1Str)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("输入无效，请输入一个有效的数字。")
		return
	}

	// 提示用户输入运算符
	fmt.Print("请输入运算符 (+, -, *, /): ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出现错误:", err)
		return
	}
	operator = strings.TrimSpace(operator)

	// 提示用户输入第二个数字
	fmt.Print("请输入第二个数字: ")
	num2Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出现错误:", err)
		return
	}
	num2Str = strings.TrimSpace(num2Str)
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("输入无效，请输入一个有效的数字。")
		return
	}

	// 使用 switch - case 语句根据运算符进行相应的计算
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("错误：除数不能为零。")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("无效的运算符，请输入 +, -, *, / 中的一个。")
		return
	}

	// 输出计算结果
	fmt.Printf("计算结果: %.2f\n", result)
}
