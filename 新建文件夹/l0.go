package main

import "fmt"
type  Bili struct {
	Up string
	FansNumber string
	LikesNumber string
	PlaybackNumber string
	Icon string
	Signature string
}

func main() {
	a:= Bili{
		Up:"老番茄",
		FansNumber:"1326w",
		LikesNumber: "6934w",
		PlaybackNumber: "14.4亿",
		Icon: "番茄",
		Signature: "没找到",
	}
	b:=Bili{Up: "李永乐",FansNumber:"666w"}
	fmt.Println(a)
	fmt.Println(b)
}
