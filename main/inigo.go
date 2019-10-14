package main
import (
	"fmt"
	"net/http"
)
// "/"的处理函数
func hello(res http.ResponseWriter,req *http.Request){
	fmt.Fprint(res,"hello,my name is Inigo Montaya")
}
func bye(res http.ResponseWriter,req *http.Request){
	fmt.Fprint(res,"bye bye")
}
func main(){
	http.HandleFunc("/",hello)
	http.HandleFunc("/bye",bye)
	http.ListenAndServe("localhost:4000",nil)
}