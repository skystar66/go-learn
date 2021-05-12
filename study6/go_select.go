package main

import (
	"fmt"
	"time"
)

func main() {

	selectCase()

}

func Chann(ch chan int, stopCh chan bool) {

	var i int
	i = 10

	for j := 0; j < 3; j++ {
		//循环发送10条消息
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func selectCase() {
	//创建channel
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)
	go Chann(ch, stopCh)
	for {
		select {
		case c = <-ch:
			fmt.Printf("C Recive %d %s", c, "from ch msg \n")
		case s := <-ch:
			fmt.Printf("S Recive %d %s", s, "from ch msg \n")
		case stop := <-stopCh:
			if stop {
				fmt.Printf("Recive true  from stopCh \n")
			} else {
				fmt.Printf("Recive false from stopCh \n")
			}
			goto end
		}
	}
	end:
	fmt.Println("结束！！！")
}
