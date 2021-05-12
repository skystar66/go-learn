package main

import (
	"fmt"
)

func main() {

	switchCase()
	//switchType()

}

func switchCase() {

	/*定义局部变量*/

	var grade string = "b"
	var score int = 90

	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch grade {
	case "A":
		fmt.Println("优秀 A")
	case "B":
		fmt.Println("良好 B")
	case "C":
		fmt.Println("及格 C")
	case "D":
		fmt.Println("不及格 D")
	default:
		fmt.Println("其它不良成绩！！！")
	}

	fmt.Printf("你的等级是：%s，你的分数是：%d \n", grade, score)
}

func switchType() {

	var x interface{}

	switch types := x.(type) {

	case nil:
		fmt.Printf(" x 的类型是 %T", types)
	case int:
		fmt.Printf(" x 的类型是 %T", types)
	case float32:
		fmt.Printf(" x 的类型是 %s", types)

	default:
		fmt.Printf(" 未知型")
	}

}
