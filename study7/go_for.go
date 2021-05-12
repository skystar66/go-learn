package main

import "fmt"

func main() {

	forTest()
	forTest2()
	for_each()
}

func forTest() {

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Printf("第 %d 次，sum：%d \n", i, sum)
	}
	fmt.Printf("sum：%d \n", sum)

}

func forTest2() {

	sum := 1
	for ; sum <= 10; {
		sum += sum
	}

	fmt.Printf("sum1：%d \n", sum)

	//另外一种写法 类似while
	for sum <= 10 {
		sum += sum
	}
	fmt.Printf("sum2：%d \n", sum)

}

func for_each() {

	strings := []string{"xuliang", "zhangyan", "zhangshuo"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	nums := [6]int{1, 2, 3, 5, 6, 9}
	for i, num := range nums {
		fmt.Printf("第 %d 位 num 的值 = %d \n", i, num)
	}

}
