package main

import (
	"fmt"
	"os"
)

func main()  {
  	os.Create("proverb.txt")
	file,_:=os.OpenFile("proverb.text",os.O_CREATE|os.O_WRONLY,os.ModePerm)
	defer file.Close()
	file.WriteString("don't communicate by sharing memory share memory by communicating")
    file,_=os.Open("proverb.text")
    bs := make([]byte,100,100)
    _,_=file.Read(bs)
    fmt.Println(string(bs))
}
