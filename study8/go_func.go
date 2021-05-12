package main

import (
	"fmt"
	"math"
)

func main() {

	var num1 = 100
	var num2 = 200

	var result = max(num1, num2)

	fmt.Printf("最大值为：%d \n", result)

	fmt.Println("---- 函数返回多个值 ----")

	str1, str2 := swap("123", "4356")

	fmt.Printf("str1: %s , str2: %s \n", str1, str2)

	fmt.Println("---- 函数使用值传递 ----")

	fmt.Printf("交换前 a 的值为 : %d\n", num1)
	fmt.Printf("交换前 b 的值为 : %d\n", num2)

	/* 通过调用函数来交换值 */
	swap_value(num1, num2)

	fmt.Printf("交换后 a 的值 : %d\n", num1)
	fmt.Printf("交换后 b 的值 : %d\n", num2)

	fmt.Println("---- 函数使用引用传递 ----")

	fmt.Printf("交换前 a 的值为 : %d\n", num1)
	fmt.Printf("交换前 b 的值为 : %d\n", num2)

	/*通过调用引用传递*/
	swap_value_ref(&num1, &num2)

	fmt.Printf("交换后 a 的值 : %d\n", num1)
	fmt.Printf("交换后 b 的值 : %d\n", num2)

	fmt.Println("---- 使用Go函数作为实参 ----")
	param_func()

	fmt.Println("---- 使用Go 闭包测试 ----")
	close_pac()

	fmt.Println("---- 使用Go 闭包参数测试 ----")
	close_pac_param()


	fmt.Println("---- 使用Go 闭包参数2测试 ----")
	close_pac_param2()

}

/**函数返回单个值*/
func max(num1, num2 int) int {

	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**函数返回多个值*/
func swap(str1, str2 string) (string, string) {

	return str1, str2
}

/**值传递*/
func swap_value(a, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp
}

/**引用传递*/
func swap_value_ref(a *int, b *int) int {

	//指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，
	//会分配独立的内存来存储一个整型数字。当变量前面有 * 标识时，
	//才等同于 & 的用法，否则会直接输出一个整型数字。

	var temp int
	temp = *a /**保持a地址上的值*/
	*a = *b   /**将b值赋值给a*/
	*b = temp /**将temp值复制给b*/
	return temp
}

/**go语言函数作为实参*/
func param_func() {

	/**声明函数变量*/
	getScore := func(score float64) float64 {
		return math.Sqrt(score)
	}

	/**使用函数*/
	fmt.Printf("获得分数 %f \n", getScore(100))

}

/**go语言闭包*/

func close_pac() {

	getSequence := getSequences()

	fmt.Printf("递增 1：%d \n", getSequence())
	fmt.Printf("递增 2：%d \n", getSequence())
	fmt.Printf("递增 3：%d \n", getSequence())

	getSequence1 := getSequences()

	fmt.Printf("递增 1：%d \n", getSequence1())
	fmt.Printf("递增 2：%d \n", getSequence1())
	fmt.Printf("递增 3：%d \n", getSequence1())

}

func getSequences() func() int {
	var i int = 0
	return func() int {
		i += 1
		return i
	}
}

/**带参数的闭包使用*/

func close_pac_param() {

	/**获取函数返回值*/
	adds := add(1, 2)
	//fmt.Printf("函数闭包返回值：%d \n",adds()
	fmt.Println(adds())
	fmt.Println(adds())
	fmt.Println(adds())

	///**提取函数返回值*/
	//i,sum:=adds()
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)

}

/**带参数的闭包使用*/

func close_pac_param2() {

	/**获取函数返回值*/
	adds := add2(1, 2,3,4)
	//fmt.Printf("函数闭包返回值：%d \n",adds()
	fmt.Println(adds(1,2,3,4))
	fmt.Println(adds(1,2,3,4))
	fmt.Println(adds(1,2,3,4))

	///**提取函数返回值*/
	//i,sum:=adds()
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)
	//fmt.Printf("提取函数返回值: i: %d  sum: %d \n",i,sum)

}

func add(x1, x2 int) func() (int, int) {

	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}

func add2(x1, x2, x3, x4 int) func(x1, x2, x3, x4 int) (int, int, int) {
	i := 0
	return func(x1, x2, x3, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}
