package main
import "fmt"
const(
	read byte=1<<iota
	write
	exec
	freeze
)
func main(){
	a:=read|write|freeze
	b:=read|freeze|exec
	c:=a&^b
	fmt.Println("%04b &^ %04b = %04b\n",a,b,c)
}