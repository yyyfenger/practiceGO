package main
import (
	"bufio"
	"fmt"
	"net"
)
func main(){
	conn,_:=net.Dial("tcp","127.0.0.1:80")//通过tcp连接
	fmt.Fprintf(conn,"GET/HTTP/1.0\r\n\r\n")
	status,_:=bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}