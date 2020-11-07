package main

import "fmt"

var (
	ch1 = make(chan bool,1)
	ch2 = make(chan bool,1)
	ch3 = make(chan bool,1)
)
func goroutine1(){
	for i:=1;i<=100;i+=2{
		<- ch1
		fmt.Println(i)
		ch2 <- true
	}
}
func goroutine2(){
	for i:=2;i<=100;i+=2{
		<- ch2
		fmt.Println(i)
		ch1 <- true
	}
	ch3 <- true
}
func main(){
	ch1 <- true
	go goroutine1()
	go goroutine2()
	<- ch3
}