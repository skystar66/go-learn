package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("1", "q", "e", 1, 2, 3, "u", "\n")
	fmt.Println("1", "q", "e", 1, 2, 3, "u")
	fmt.Printf("ab %d %d %d cd", 1, 2, 3)

	fmt.Println("《《===========Go语言基础数据类型============》》")

	var numdouble = 3.5
	var num = 3

	fmt.Println("数值类型", numdouble, num)

	var avliable bool //声明变量
	valide := true    //简短声明
	avliable = true   //赋值变量

	fmt.Println("输出bool 类型", avliable, valide)

	fmt.Println("《《===========Go语言字符串声明变量及去除空格============》》")

	str := "这是一个 简短的字符串声明x变量"
	fmt.Println("原字符串")
	fmt.Println(str)

	//去除空格
	str = strings.Replace(str, " ", "", -1)
	//去除指定符号x
	str = strings.Replace(str, "x", "", -1)
	fmt.Println("去除空格及换行符后的字符串")
	fmt.Println(str)

}
