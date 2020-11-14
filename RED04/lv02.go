package main

import (
	"fmt"
	"time"
)

func main(){
	t1 := time.NewTicker(1*time.Minute)
	for  {
		<- t1.C
		fmt.Println("芜湖! 起飞")
	}
	func(){
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			fmt.Printf("没有困难的工作，只有勇敢的打工人！")
		}
	}()
	func(){
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 8, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			fmt.Printf("早安，打工人！")
		}
	}()
}

