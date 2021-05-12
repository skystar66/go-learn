package main

import (
	"fmt"
)

func main() {
	//mathFunc()

	//otherMath()

	//valueRef()

	changeValueRef()
}

func mathFunc() {

	var a = 21
	var b = 10
	var c = a + b

	fmt.Println("第一行 - c 的值为 %d\\n", c)
	c = a - b
	fmt.Println("第二行 - c 的值为 %d\\n", c)
	c = a * b
	fmt.Println("第三行 - c 的值为 %d\\n", c)
	c = a / b
	fmt.Println("第四行 - c 的值为 %d\\n", c)
	c = a % b
	fmt.Println("第五行 - c 的值为 %d\\n", c)
	a++
	fmt.Println("第六行 - a 的值为 %d\\n", c)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

func setValue() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

}

func otherMath() {
	var a int =4
	var b int32
	var c float32
	var ptr *int

	/*运算符实例*/
	fmt.Printf("第1行 - a 变量类型为 = %T\n",a)
	fmt.Printf("第2行 - b 变量类型为 = %T\n",b)
	fmt.Printf("第3行 - c 变量类型为 = %T\n",c)
	/*& 和 * 运算符实例*/
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a的值为 %d \n",a)
	fmt.Printf("ptr的值为 %d \n",*ptr)

	a=100
	fmt.Printf("a的值为 %d \n",a)
	fmt.Printf("ptr的值为 %d \n",*ptr)

}

func valueRef() {

	var a int =4
	var b *int  //声明一个指针变量

	b= &a //将a的地址值 赋值给b
	println("a的值为", a)   // 4
	println("*b为", *b)  // 4
	println("b为", b)   // 824633794744


}


func changeValueRef() {

	var a int =4
	var b *int  //声明一个指针变量
	b= &a //将a的地址值 赋值给b
	println("a的值为", a)   // 4
	println("*b为", *b)  // 4
	println("b为", b)   // 824633794744

	a= 100
	println("a的值为", a)   // 100
	println("*b为", *b)  // 100
	println("b为", b)   // 824633794744



}





















