package main

import (
	"fmt"
)

func main() {

	range_test()
}

func range_test() {

	arrays := []int{1, 2, 3, 4, 5}

	for index, val := range arrays {
		fmt.Printf("index:%d,value:%d \n", index, val)
	}

	maps := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range maps {
		fmt.Printf("key:%s,value:%s \n", k, v)
	}

}
