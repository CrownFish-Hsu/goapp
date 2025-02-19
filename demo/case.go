package demo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowGradeLevel() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入成绩（0 - 100）：")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("input error:", err)
	}

	input = strings.TrimSpace(input)
	fmt.Printf("input=%s, type=%T", input, input)

	score, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("parse float error:", err)
	}

	var grade string
	switch {
	case score >= 80:
		grade = "A"
	case score >= 60:
		grade = "B"
	default:
		grade = "C"
	}

	fmt.Printf("score grade is %s\n", grade)
}
