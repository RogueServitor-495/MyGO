package mylib

import "fmt"

// 首字母大写的变量为公共变量，可在其他包中访问
var Exportable int = 42

// 首字母小写的变量为私有变量，只能在包内部访问
var PackName string = "Chenph"

func Myfunc(){
	fmt.Println(Exportable)
	fmt.Println(PackName)
}