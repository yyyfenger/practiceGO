package main
import (
	"fmt"
	"time"
)
func printCount(c chan int){//传递进一个整形通道
	num:=0
	for num>=0{
		num=<-c//等待函数值进入通道
		fmt.Print(num," ")
	}
}
func main(){
	c:=make(chan int)//创建一个通道
	a:=[]int{8,7,6,5,3,0,9,-1}
	go printCount(c)
	for _,v:=range a{
		c<-v//将整形变量传递进入通道
	}
	time.Sleep(time.Millisecond*1)//全部质量零结束前暂停
	fmt.Println("End of main")
}