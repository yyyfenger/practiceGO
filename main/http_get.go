package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main(){
	resp,_:=http.Get("http://example.com/")//发出获取http请求
	body,_:=ioutil.ReadAll(resp.Body)//从响应中读取body
	fmt.Println(string(body))
	resp.Body.Close()//关闭连接
}