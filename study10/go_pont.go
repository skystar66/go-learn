package main

import (
	"fmt"
)

const (
	array_count = 3
)

func main() {

	point_test1()

	point_array()

	point_array_error()

	point_array2()

	point_to_point2()

}

func point_test1() {

	//声明一个指针变量
	var a *int

	//声明一个实际变量

	var b int = 20

	//将实际变量的内存存储地址赋予给指针变量的内存存储地址
	a = &b

	//打印实际变量b的内存存储地址
	fmt.Printf("b的内存存储地址:%d \n", &b)

	//打印指针变量a的内存存储地址
	fmt.Printf("a的内存存储地址:%d \n", a)

	//打印指针变量a的实际值
	fmt.Printf("*a的实际值:%d \n", *a)

}

func point_array() {
	//声明一个指针数组
	var pointArray [array_count]*int
	//声明一个实际数组
	realArray := []int{100, 200, 300}
	//循环为指针数组赋值
	for i := 0; i < len(realArray); i++ {
		//将实际数组的内存存储值赋值给指针数组的内存存储值
		pointArray[i] = &realArray[i]
	}
	//打印数组的实际值
	for k, val := range pointArray {
		fmt.Printf("pointArray 索引:%d,内存存储值:%d, 实际值:%d \n", k, val, *val)

	}
}

func point_array_error() {

	/*从结果中我们发现内存地址都一样，而且值也一样。怎么回事？
	  这个问题是range循环的实现逻辑引起的。跟for循环不一样的地方在于range循环中的x变量是临时变量。range循环只是将值拷贝到x变量中。因此内存地址都是一样的。*/
	number := [array_count]int{5, 6, 7}
	var ptrs [array_count]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i, x := range &number {
		ptrs[i] = &x
		fmt.Printf("number的内存存储值为:%d ,实际值为:%d \n", &x,x)
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}



func point_array2() {


	/*flying81621 的 for range 方法也只是用于遍历数组，并没用使用 range 方法的临时变量。真正赋值的地方只有一次，for 的 body 内 ，青云刀歌真正原因在于，
	x临时变量仅被声明一次，此后都是将迭代 number 出的值赋值给 x ， x 变量的内存地址始终未变，这样再将 x 的地址发送给 ptrs 数组，自然也是相同的。*/
	number := [array_count]int{5, 6, 7}
	var ptrs [array_count]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i, x := range &number {
		temp:=x //此时将变量值复制给临时变量，内存存储值没有改变
		ptrs[i] = &temp
		fmt.Printf("number的内存存储值为:%d ,实际值为:%d \n", &temp,temp)
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}

func point_to_point2() {


	//声明一个实际变量
	var a int =3000

	//声明一个指针变量
	var b *int

	//声明一个指向指针的指针变量
	var c **int
	//将a的内存存储值赋值给指针变量b
	b=&a

	//将b的内存地址存储值 赋值给 c
	c=&b

	fmt.Printf("a的值为：%d \n",a)
	fmt.Printf("b的值为：%d \n",*b)
	fmt.Printf("c的值为：%d \n",**c)


}



























