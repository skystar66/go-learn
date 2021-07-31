package main

import (
	"fmt"
	"time"
)

func main() {

	//go say("word")
	//say("hello")

	test_chann()

	test_make_more_channel()

}

func say(str string) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把sum发送到通道c
}

func test_chann() {

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从通道c接收
	fmt.Println(x, y, x+y)

}

func test_make_more_channel() {

	//定义一个可以存储大小为2的缓冲区通道
	ch:=make(chan int,2)

	ch<-1
	ch<-2

	//获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}






