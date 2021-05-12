package main

import "fmt"

func main() {

	goto_test()
}

func goto_test() {

	/*定义局部变量*/

	var num = 10

LOOP:
	for num <= 20 {
		if num == 15 {
			num += 1
			fmt.Printf("执行 goto 语句 ，num 的值为 %d \n", num)
			goto LOOP
		}
		fmt.Printf("num 的值为 %d \n", num)
		num++
	}

}
