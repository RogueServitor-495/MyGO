package main

import "fmt"

func main(){
	fmt.Println("this is a file for teaching common data structures in golang")

	/*
	Array数组：长度固定，不可拓展
	*/
	fmt.Println("=============Arrays:=============")
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Println(arr1)
	
	arr1[0] = 4
	arr1[2] = 6
	fmt.Println(arr1)

	arr2 := [3]int{4,5,6}
	fmt.Printf("type:%T\n",arr2)
	fmt.Println(arr2)

	/*
	Map结构，键值对
	*/
	var mp1 map[string]int

}