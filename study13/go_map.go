package main

import (
	"fmt"
)

func main() {
	create_map()

	map_delete()

}
func create_map() {

	//创建声明一个map类型
	var testMap map[string]string
	testMap = make(map[string]string)
	//为map赋值
	testMap["beijing"] = "北京"
	testMap["shanghai"] = "上海"
	testMap["guangzhou"] = "广州"
	testMap["shenzhen"] = "深圳"
	testMap["changchun"] = "长春"
	//使用range 循环，通过键输出值
	for s := range testMap {
		fmt.Printf("key:%s,value:%s \n", s, testMap[s])
	}
	//查看元素是否存在
	captial, ok := testMap["haerbin"]

	fmt.Println(ok)
	fmt.Println("captial:" + captial)
	if ok {
		fmt.Println("存在哈尔滨城市")
	} else {
		fmt.Println("不存在哈尔滨城市")

	}
}

func map_delete() {
	var testMap = map[string]string{"beijing": "北京", "shanghai": "上海", "changchun": "长春"}

	for s := range testMap {
		fmt.Printf("key:%s,value:%s \n", s, testMap[s])
	}

	//删除元素
	delete(testMap, "beijing")

	fmt.Println("删除后元素遍历")

	for s := range testMap {
		fmt.Printf("key:%s,value:%s \n", s, testMap[s])
	}
}
