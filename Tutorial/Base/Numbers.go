package main

import "fmt"

func tuto(){
	fmt.Println("This file teaches types of numbers")

	var f32 float32 = 0.1
	var f64 float64 = 0.1
	
	fmt.Println("test")
	var i8 int8 = 127
	var sum int8 = i8+2
	
	fmt.Printf("sum的数据类型：%T，值为：%d\n" , sum,sum)

	println(f32 - float32(f64))
}

func main() {
	tuto()
}