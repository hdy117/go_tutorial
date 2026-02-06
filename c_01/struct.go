package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
	Tele  string
}

// 获取姓名
func (obj *Person) GetName() string {
	return obj.Name
}

// 获取年龄
func (obj *Person) GetAge() int {
	return obj.Age
}

// 获取邮箱
func (obj *Person) GetEmail() string {
	return obj.Email
}

// 获取电话
func (obj *Person) GetTele() string {
	return obj.Tele
}

// 显示信息， 实现IShow接口
func (obj *Person) ShowInfo() {
	fmt.Printf("addr:%p.\n", obj)
	fmt.Printf("Name: %s, Age: %d, Email: %s, Tele: %s", obj.Name, obj.Age, obj.Email, obj.Tele)
}
