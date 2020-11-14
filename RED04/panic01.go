package main
import "fmt"
func Demo1(){
	defer func() {
		panic("first")
	}()
	defer func() {
		panic("second")
	}()
	fmt.Println("main function222")
}
func main()  {
	go Demo1()
	Demo2()
}
func Demo2(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("first defer panic")
	}()
	defer func() {
		panic("second defer panic")
	}()
	panic("main body panic")
}