package main

import "fmt"

func main() {
	test_error()
}

type numMath struct {
	chushu    int
	beichushu int
}

//实现error接口
func (num *numMath) Error() string {
	errorStrFormat := "Cannot proceed,this beichushu is zero. chushu:%d,beichushu:0"
	return fmt.Sprintf(errorStrFormat, num.chushu)
}

//定义函数
func divied(chushu int, beichushu int) (result int, string2 string) {
	if beichushu == 0 {
		num := numMath{beichushu: beichushu, chushu: chushu}
		return 0, num.Error()
	} else {
		return chushu / beichushu, ""
	}
}

func test_error() {

	//正常情况
	result, error := divied(100, 1)
	fmt.Printf("result:%d,str:%s \n", result, error)
	//错误情况
	restorer, error2 := divied(100, 0)
	fmt.Printf("result:%d,str:%s \n", restorer, error2)

}
