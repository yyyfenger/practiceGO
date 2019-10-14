package main
import "fmt"
func main(){
	x:=make([]int,3,5)
	for i:=0;i<8;i++{
		x=append(x,i)
	}
	fmt.Println(x)
}