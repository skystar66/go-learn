package main

import (
	"fmt"
)

func main() {


	array_test1()

	array_test2()

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
	balance:=[5] float32{1000.0,2.0,3.4,7.0,50.0}

	//输出数组元素
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f \n",i,balance[i])
	}

	balance2:=[...]float32{1000.0,2.0,3.4,7.0,50.0}
	//输出数组元素
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f \n",j,balance2[j])
	}

	//将索引1 3 的元素进行初始化
	balance3:=[5]float32{1:2.0,3:30.0}
	//输出数组元素
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f \n",k,balance3[k])
	}


}

