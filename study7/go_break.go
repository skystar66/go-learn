package main

import "fmt"

func main() {

	//for_break()

	//break_re()

	//for_continue()

	for_continue_label()
}

func for_break() {

	var num = 1

	for num < 20 {

		fmt.Printf("num 的值为： %d \n", num)

		num++

		if num > 15 {
			/* 使用 break 语句跳出循环 */
			fmt.Printf("num 的值为： %d 跳出循环 \n", num)
			break
		}
	}

}

func break_re() {

	fmt.Println("-------不使用break label -------")

	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d \n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d \n", i2)
			break
		}
	}

	fmt.Println("-------使用break label -------")
label:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d \n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d \n", i2)
			break label
		}
	}
}

func for_continue() {

	/*定义局部变量*/
	var a = 15
	for a < 20 {
		if a == 15 {
			a = a + 1
			fmt.Println("触发contnue 语句，执行下次循环！！！")
			continue
		}
		fmt.Printf("a 的值为 %d \n", a)
		a++
	}
}

