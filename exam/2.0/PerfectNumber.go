package main

import (
	"fmt"
	"time"
)
func main(){
	start:=time.Now()
	go getnumber()
	go run()
	t:=time.Now().Sub(start)

	fmt.Println("t=",t)//
}
func getnumber(){
	for i:=1;i<123456;i++ {
			num:=addfactor(i)
			if num==i {
			fmt.Printf("%d  ",i)
		}
	}
}

func addfactor(i int)int {
	var num=0
	for j:=1;j<i;j++ {

		if i%j==0{num+=j}
	}
	return num
}
func run() {
	for{}
}
