package main

import (

	"fmt"
)

func main()  {
	//map 的声明
	var m map[string]string
	n:=map[string]int {"one":1,"two":2}
	//map 的初始化
	m =map[string]string{"one":"first","two":"second"}
	fmt.Println("one:",m["one"])
	fmt.Println("one：",n["one"])
	m["one"]="first1(1)"
	//map的覆盖
	fmt.Println("one1",m["one"])
	//map 的删除delete
	//delete(m,"two")
	fmt.Println(m)
	//range map
	for key ,value :=range n{
		fmt.Println(key ,value)
	}
}
