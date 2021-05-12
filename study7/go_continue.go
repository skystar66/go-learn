package main

import "fmt"

func main() {

	for_continue_label()
}


func for_continue_label() {

	fmt.Println("----- 不使用 continue 标识 label -----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i 为：%d \n", i)
		for i2 := 11; i2 < 15; i2++ {
			fmt.Printf("i2 为：%d \n", i2)
			continue
		}
	}
	fmt.Println("----- 使用 continue 标识 label -----")

con_label:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i 为：%d \n", i)
		for i2 := 11; i2 < 15; i2++ {
			fmt.Printf("i2 为：%d \n", i2)
			continue con_label
		}
	}

}