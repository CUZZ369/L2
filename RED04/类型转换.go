package main

import (
	"fmt"
	"strconv"
	"strings"
)
//func Parse1(s string,base int ,bitesize int )(i int64 ,err error)
//func Parse2(s,string,base int ,bitesize int )(i uint64,err error)
//func Atoi(s string)(i int ,err error)
func main()  {
	fmt.Println(strconv.ParseInt("A",16,0))
	fmt.Println(strconv.ParseUint("16",8,0))
	fmt.Println(strconv.Itoa(22))

	fmt.Println(strconv.ParseFloat("3.1415926",32))
	fmt.Println(strconv.ParseFloat("3.1415926",64))

	fmt.Println(strconv.FormatFloat(3.1415926,'E',+1,32))
	fmt.Println(strconv.FormatFloat(3.1415626,'e',-1,64))

	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.FormatBool(true))

 	s:="abcdefg"
	fmt.Println(strings.Index(s,"e"))
	fmt.Println(strings.Index(s,"a"))
	fmt.Println(strings.Index(s,"h"))

	fmt.Println(strings.Contains("TigerwolfC", "wolf"))
	fmt.Println(strings.Contains("TigerwolfC", "bar"))
	fmt.Println(strings.Contains("TigerwolfC", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("我是中国人", "中国"))
}