package main

import (
	"fmt"
)

func main() {
	test_Interface()

	test_return_interface()

	test_param_interface()

	test_Model()

	test_param_ref_interface()

	test_update_from_interface()

	test_group_interface()

}

//声明接口
type Phone interface {
	call()
}

type NioKiaPhone struct {
}

func (NioKiaPhone NioKiaPhone) call() {
	fmt.Println("I am NioKia,I can call you!!")
}

type IPhone struct {
}

func (IPhone IPhone) call() {

	fmt.Println("I am phone,I can call you!!")
}

func test_Interface() {

	phone := new(IPhone)
	phone.call()

	nioKiaPhone := new(NioKiaPhone)

	nioKiaPhone.call()

}

//声明一个返回结果的接口
type Man interface {
	name() string
	age() int
}

type Nike struct {
}

type Lili struct {
}

func (nike2 Nike) name() string {

	return "nike"
}

func (nike2 Nike) age() int {
	return 13
}

func (lili Lili) name() string {

	return "lili"
}
func (lili Lili) age() int {
	return 14
}

func test_return_interface() {

	var man Man
	man = new(Nike)
	fmt.Printf("你好，我的名字叫：%s 年龄:%d \n", man.name(), man.age())
	man = new(Lili)
	fmt.Printf("你好，我的名字叫：%s 年龄:%d \n", man.name(), man.age())
}

//声明一个传参的接口c
type Men interface {
	name(name string) string
	age() int
}

type Marry struct {
}

func (marry Marry) name(name string) string {
	return name
}

func (marry Marry) age() int {
	return 14
}

func test_param_interface() {
	name := "marray"
	var men Men
	men = new(Marry)
	fmt.Printf("name：%s \n", men.name(name))
	fmt.Printf("age：%d \n", men.age)
}

//接口案例
type Detect interface {
	create_model() Model
}

//定义结构体 创建属性
type Model struct {
	name string
	age  int
}

func (model Model) create_model() Model {
	model.name = "xl"
	model.age = 25
	return model
}

func test_Model() {

	var detect Detect
	detect = new(Model)
	model := detect.create_model()
	fmt.Printf("name:%s age: %d \n", model.name, model.age)
}

//将接口作为参数
type UserInfo interface {
	usernName() string
	userAge() int
}

type XL struct {
	name string
	age  int
}

type ZY struct {
	name string
	age  int
}

func (xl XL) usernName() string {
	xl = XL{age: 12, name: "xl"}
	return xl.name
}

func (xl XL) userAge() int {
	xl = XL{age: 12, name: "xl"}
	return xl.age
}

func (zy ZY) usernName() string {
	zy = ZY{age: 14, name: "zy"}

	return zy.name
}

func (zy ZY) userAge() int {
	zy = ZY{age: 14, name: "zy"}
	return zy.age
}

func schedule(info UserInfo) {
	fmt.Printf("我的名字：%s 年龄:%d \n", info.usernName(), info.userAge())
}

func test_param_ref_interface() {

	var info UserInfo
	info = new(XL)
	schedule(info)

	info = new(ZY)
	schedule(info)

}

//通过接口方法修改属性
type fruit interface {
	geName() string
	setName(name string)
}
type apple struct {
	name string
}

func (a *apple) getName() string  {
	return a.name
}
func (a *apple) setName(name string)  {
	 a.name = name
}
func test_update_from_interface() {

	apl:=apple{
		name: "红富士",
	}
	fmt.Printf("苹果品种：%s \n",apl.getName())
	apl.setName("烟台栖霞苹果")
	fmt.Printf("苹果品种：%s \n",apl.getName())

}


//组合接口的实现
type reader interface {
	reading() string
}
type writer interface {
	writing() string
}

type rw interface {
	reader
	writer
}

type mourse struct {

}

func (m mourse) reading() string  {
	return "正在读书......"
}

func (m *mourse) writing() string {
	return "正在写作....."
}

func test_group_interface() {
	var rw1 rw
	rw1 = &mourse{}
	fmt.Println(rw1.reading())
	fmt.Println(rw1.writing())
}




























