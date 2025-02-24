package demo

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Score string `json:"score"`
}

type Class struct {
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"`
}

func TestStudent() {
	// 使用 new 创建指针并初始化零值
	s := new(Student)
	s.Name = "Allen"
	s.Age = 18
	s.Score = "100"

	fmt.Printf("通过 new 创建的学生:Name=%s,Age=%d,Score=%s, Type=%T\n", s.Name, s.Age, s.Score, s)

	fmt.Println()

	// 使用结构体字面量直接创建实例（值类型）
	s1 := Student{
		Name:  "Iverson",
		Age:   30,
		Score: "98",
	}

	s1.Age = 40
	s1.Score = "70"

	fmt.Printf("通过字面量创建的学生:Name=%s,Age=%d,Score=%s, Type=%T\n", s1.Name, s1.Age, s1.Score, s1)
}

func TestCalss() {
	class := Class{
		"S-Class",
		[]Student{
			{"Alice", 18, "90"},
			{"Bob", 19, "88"},
			{"Charlie", 17, "85"},
		},
	}

	class.Students = append(class.Students, Student{"David", 20, "80"})
	fmt.Printf("班级名称:%s\n", class.ClassName)
	for idx, student := range class.Students {
		fmt.Printf("%d.学生姓名:%s,年龄:%d,成绩:%s\n", idx+1, student.Name, student.Age, student.Score)
	}

	jsonData, _ := json.Marshal(class)
	fmt.Println(string(jsonData))
}
