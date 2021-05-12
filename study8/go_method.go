package main

import "fmt"

/*定义结构体*/
type Circle struct {
	name string
}

func main() {
	//var c1 Circle
	//c1.name="xuliang"
	//fmt.Println("你的姓名=",c1.getAllName())

	var c Circle
	fmt.Println(c.name)
	c.name = "xl"
	fmt.Println(c.getAllName())
	c.changeRadius("sss")
	fmt.Println(c.name)
	change(&c, "zzssd")
	fmt.Println(c.name)

}


//该 method 属于 Circle 类型对象中的方法
func (c Circle) getAllName() string {
	return "xl001"+c.name+"----xl002"+c.name
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(name string)  {
	c.name = name
}

// 以下操作将不生效
//func (c Circle) changeRadius(radius float64)  {
//   c.radius = radius
//}

// 引用类型要想改变值需要传指针
func change(c *Circle, name string)  {
	c.name = name
}











