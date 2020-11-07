package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {

		go func() {
			<-over
			fmt.Println(i)
		}()
		if i != 10 {
			over <- true
		}
	}

	fmt.Println("over!!!")
}