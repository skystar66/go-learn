package main

import (
	"fmt"
	"unsafe"
)

func main() {

	slice1()

	fmt.Println("以下为 切片的 append and copy ")
	slice_append_copy()

	fmt.Println("以下为 合并数组 ")
	mergeArray()

	fmt.Println("以下为 切片的append and copy")
	appendAndCopy2()

}

func slice1() {

	//创建切片
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	printSlice(numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Printf("numbers[1:4]:%v\n", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Printf("numbers[:3]:%v\n", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Printf("numbers[4:]:%v\n", numbers[4:])
	//使用 make函数创建切片
	numbers1 := make([]int, 3, 5)
	printSlice(numbers1)
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	numbers2 := numbers1[0:2]
	fmt.Printf("make numbers1[0:2]:%v\n", numbers2)
	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	numbers3 := numbers1[2:5]
	fmt.Printf("make numbers1[2:5]:%v\n", numbers3)

}

func slice_append_copy() {

	var numbers []int
	printSlice(numbers)
	/**追加空切片 赋值0*/
	numbers = append(numbers, 0)
	printSlice(numbers)
	/**追加单个值 */
	numbers = append(numbers, 1)
	printSlice(numbers)
	/**追加多个值*/
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/**扩容切片容量为之前切片的2倍*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/**将numbers的内容copy 到numjbers1中*/
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func mergeArray() {
	num1 := []int{1, 2, 3}
	num2 := []int{4, 5, 6}
	num3 := []int{7, 8, 9}
	num4 := append(append(num1, num2...), num3...)
	printSlice(num4)

}

func appendAndCopy2() {
	array1 := []int{1, 2}
	params := []int{2, 3, 4}
	fmt.Printf("数组的长度为：%d 参数的长度为:%d 和为:%d 数组的cap 为:%d  参数的cap 为:%d \n",
		len(array1), len(params), len(array1)+3, cap(array1), cap(params))
	array1 = append(array1, params...)
	printSlice(array1)

}

/*1，当把 slice 作为参数，本身传递的是值，但其内容就 byte* array，实际传递的是引用，所以可以在函数内部修改
2，但如果对 slice 本身做 append，而且导致 slice 进行了扩容，实际扩容的是函数内复制的一份切片，对于函数外面的切片没有变化。
*/
func slice_test() {

	slice_test := []int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(slice_test))
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_value(slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_value_ref(&slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	fmt.Println(unsafe.Sizeof(slice_test))

}

/**slice自带引用*/
func slice_value(sliceVal []int) {
	sliceVal[1] = 100 //函数外的slice 改变
	sliceVal = append(sliceVal, 6) //函数外的slice 不改变
	fmt.Printf("slice_value:%#v,%#v,%#v\n", slice_test, len(sliceVal), cap(sliceVal))
}

/**修改函数外的切片slice*/
func slice_value_ref(sliceVal *[]int) {
	*sliceVal = append(*sliceVal, 6) //函数外的slice 改变
	fmt.Printf("slice_value_ref:%#v,%#v,%#v\n", *sliceVal, len(*sliceVal), cap(*sliceVal))
}

func printSlice(numbers []int) {
	fmt.Printf("length:%d,cap:%d slice:%v\n", len(numbers), cap(numbers), numbers)
}
