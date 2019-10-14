package main
import (
	"fmt"
	"time"
)
func count(){
	for i:=0;i<5;i++{
		fmt.Println(i)
		time.Sleep(time.Millisecond*1)
	}
}
func main(){
	go count()//运行goroutine函数
	time.Sleep(time.Millisecond*2)
	fmt.Printf("hello world")
	time.Sleep(time.Millisecond*5)
}