package main

import "fmt"

const (

	/*在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
	*/
	a = 1
	b
	c
	d

)

/*iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。
*/
const (
	i = iota
	j = iota
	x = iota
)
const xx = iota
const yy = iota



func main() {

	fmt.Println(a)
	// b、c、d没有初始化，使用上一行(即a)的值
	fmt.Println(b)   // 输出1
	fmt.Println(c)   // 输出1
	fmt.Println(d)   // 输出1


	println(i, j, x, xx, yy)
	// 输出是 0 1 2 0 0

}
