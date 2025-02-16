package main

import "fmt"


var globalage = 15


func main() {
	fmt.Println("声明变量")
	/*
	这里是多行注释
	- var 声明go变量，基本格式 var {name} {type} 
	- 变量名 由数字、字母、下划线组成，首字符不能为数字
	- 支持自动类型推导
	*/

	// String类型	初始值为空  
	var name string = "test"	
	fmt.Println("声明了一个String变量name:" + name)

	var _name1 = "test1"
	fmt.Println(_name1);


	// int 类型
	var age = 15
	println(age)

	/*
	自动类型推导：象牙符号':='(仅在函数内使用)
	*/

	myage := 15
	print(myage)

	myname := "jack"
	println(myname)

	/*
	一次定义多个变量
	1. var v1,v2,...,vn type
	2. var (
		v1 type1
		v2 type2
		v3 type3
		)
	*/

	// var v1,v2,v3 int = 1,2,3

	v4,v5,v6 := 1,2,"test"

	fmt.Println(v4,v5,v6)

	// 打印数据类型 printf %T

	fmt.Printf("数据类型：%T",myname)

}