package demo

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

func TestVar() {
	TestVarType()

	TestIntVarConvert()

	TestIntFloatVarConvert()

	TestIntStringVarConvert()

	TestBigIntStringVarConvert()

	TestReflectVarConvert()
}

func TestVarType() {
	// 声明整数类型变量
	var num1 int = 10
	var num2 int8 = 20
	var num3 int16 = 30
	var num4 int32 = 40
	var num5 int64 = 50

	// 声明无符号整数类型变量
	var unum1 uint = 10
	var unum2 uint8 = 20
	var unum3 uint16 = 30
	var unum4 uint32 = 40
	var unum5 uint64 = 50

	// 声明浮点数类型变量
	var float1 float32 = 10
	var float2 float64 = 20

	// 声明布尔类型变量
	var isTrue bool = true
	var isFalse bool = false

	// 声明字符串类型变量
	var str string = "Hello, Go!"

	fmt.Printf("num1: %d, type: %T\n", num1, num1)
	fmt.Printf("num2: %d, type: %T\n", num2, num2)
	fmt.Printf("num2: %d, type: %T\n", num3, num3)
	fmt.Printf("num2: %d, type: %T\n", num4, num4)
	fmt.Printf("num2: %d, type: %T\n", num5, num5)
	fmt.Println()

	fmt.Printf("num6: %d, type: %T\n", unum1, unum1)
	fmt.Printf("num6: %d, type: %T\n", unum2, unum2)
	fmt.Printf("num6: %d, type: %T\n", unum3, unum3)
	fmt.Printf("num6: %d, type: %T\n", unum4, unum4)
	fmt.Printf("num6: %d, type: %T\n", unum5, unum5)
	fmt.Println()

	fmt.Printf("float1: %f, type: %T\n", float1, float1)
	fmt.Printf("float1: %f, type: %T\n", float2, float2)
	fmt.Println()

	fmt.Printf("isTrue: %v, type: %T\n", isTrue, isTrue)
	fmt.Printf("isTrue: %v, type: %T\n", isFalse, isFalse)
	fmt.Println()

	fmt.Printf("str: %s, type: %T\n", str, str)
	fmt.Println()
}

func TestIntVarConvert() {
	// 声明整数类型变量
	var num1 int = 10
	var num8 int8 = int8(num1)
	var num16 int16 = int16(num8)
	var num0 int = int(num16)

	fmt.Printf("num1: %d, type: %T\n", num1, num1)
	fmt.Printf("num2: %d, type: %T\n", num8, num8)
	fmt.Printf("num2: %d, type: %T\n", num16, num16)
	fmt.Printf("num2: %d, type: %T\n", num0, num0)
	fmt.Println()
}

func TestIntFloatVarConvert() {
	// 整数转换为浮点数
	var num int = 200
	var float1 float64 = float64(num)
	fmt.Printf("num3: %d, type: %T\n", num, num)
	fmt.Printf("float1: %f, type: %T\n", float1, float1)
	fmt.Println()

	// 浮点数转换为整数
	var float2 float64 = 3.14
	var num4 int = int(float2)
	fmt.Printf("float2: %f, type: %T\n", float2, float2)
	fmt.Printf("num4: %d, type: %T\n", num4, num4)
	fmt.Println()
}

func TestIntStringVarConvert() {
	var numStr string = "123"
	num5, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("转换出错:", err)
	} else {
		fmt.Printf("numStr: %s, type: %T\n", numStr, numStr)
		fmt.Printf("num5: %d, type: %T\n", num5, num5)
	}
	fmt.Println()

	num6 := 456
	numStr2 := strconv.Itoa(num6)
	fmt.Printf("num6: %d, type: %T\n", num6, num6)
	fmt.Printf("numStr2: %s, type: %T\n", numStr2, numStr2)
	fmt.Println()

	// 字符串转换为浮点数
	strFloat := "3.14"
	floatNum, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Printf("字符串转换为浮点数出错: %v\n", err)
	} else {
		fmt.Printf("转换后的浮点数: %f\n", floatNum)
	}
	fmt.Println()
}

func TestBigIntStringVarConvert() {
	// 字符串转换为大整数
	numStr := "12345678901234567890"
	bigInt := new(big.Int)
	_, success := bigInt.SetString(numStr, 10)
	if success {
		fmt.Printf("转换后的大整数: %s, type: %T\n", bigInt.String(), bigInt)
	} else {
		fmt.Println("字符串转换为大整数失败")
	}
	fmt.Println()

	// 大整数转换为字符串
	str := bigInt.String()
	fmt.Printf("大整数转换后的字符串: %s, type: %T\n", str, str)
	fmt.Println()
}

func TestReflectVarConvert() {
	num := 10
	numValue := reflect.ValueOf(num)
	convertType := reflect.TypeOf(int64(0))
	// 转换为 int64 类型
	if numValue.Kind() == reflect.Int {
		newNum := numValue.Convert(convertType).Int()
		fmt.Printf("转换后的 int64 类型值: %d, %T\n", newNum, newNum)
	}

	fmt.Println()
}
