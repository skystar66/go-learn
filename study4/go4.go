package main

import "fmt"
import "unsafe"

const (

	a="abc"
	b= len(a)
	c= unsafe.Sizeof(a) //字符串大小字节数

)



func main() {

	//refType()
	//
	//a,b, c :=returnNums()
	//fmt.Println(a,b,c)


	fmt.Println("=====常量枚举函数=====")
	fmt.Println(a,b,c)

	fmt.Println("=====常量使用=====")
	costTest();






}

func refType() {
	var num = 123
	num2 := num

	fmt.Println("num2:", num2, ",num1:", num)

	num = 234

	fmt.Println("num2:", num2, ",num1:", num)

	var str1 = "123"
	var str2 = str1

	fmt.Println("str1:", str1, ",str2:", str2)

	str1 = "4567"
	fmt.Println("str1:", str1, ",str2:", str2)

}

/*
一个可以返回多个值的函数
*/
func returnNums() (int, int, string) {
	a, b, c := 1, 2, "123"

	return a, b, c

}

/**常量的使用*/
func costTest() {

	const LENGTH int = 10
	const WIDTH int= 5
	var area int
	//多常量赋值
	const a, b, c = 1, false, "123"

	area=LENGTH*WIDTH
	fmt.Printf("面积为：%d",area)
	fmt.Println()
	fmt.Println(a,b,c)
}













