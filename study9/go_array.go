package main

import (
	"fmt"
)

func main() {

	array_test1()

	array_test2()

	array_test3()

	array_test4()


	array_test5()

}

func array_test1() {
	var n [10]int /*n是一个长度为10的数组*/
	var i, j int

	/*为数组n初始化元素*/
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	/*输出每个数组元素的值*/
	for j = 0; j < 10; j++ {
		fmt.Printf("out array[%d] value:%d \n", j, n[j])
	}

}

func array_test2() {

	var i, j, k int

	//声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//输出数组元素
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f \n", i, balance[i])
	}

	/*动态声明数组元素大小*/
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//输出数组元素
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f \n", j, balance2[j])
	}

	//将索引 1 3 的元素进行初始化
	balance3 := [5]float32{1: 2.0, 3: 30.0}
	//输出数组元素
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f \n", k, balance3[k])
	}
}

//二维数组
func array_test3() {

	//Step 1:创建数组
	values := [][]int{}
	//Step 2 使用append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	//Step 3：显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])

	fmt.Println("Row 2")

	fmt.Println(values[1])

	//Step 4 ： 访问第一个元素
	fmt.Println("第一个元素为：")

	fmt.Println(values[0][0])

}

//初始化二维数组
func array_test4() {

	// 创建二维数组
	sites := [2][2]string{}

	//向二维数组添加元素
	sites[0][0] = "java"
	sites[0][1] = "c++"

	sites[1][0] = "python"
	sites[1][1] = "golang"

	//显示结果
	fmt.Println(sites)
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	var i, j int

	/* 输出数组元素 */
	for  i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
}

func array_test5() {

	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}

}



















