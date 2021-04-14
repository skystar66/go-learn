package main

import "fmt"
import "../mymath"

func main() {

	fmt.Println("Google"+"Runoob")
	formatStr()

	mathClass.Add(1,2)
	mathClass.Sub(2,1)

}

func formatStr() {
	var stockcode=123
	var datestyr = "2021-04-12"
	var url ="code=%d & date=%s"
	var target_url=fmt.Sprintf(url,stockcode,datestyr)

	fmt.Println("格式化后的Url："+target_url)
}



