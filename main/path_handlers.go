package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)
func main(){
	pr:=newPathResolver()
	pr.add("GET/hello",hello)
	pr.add("*/goodbye/*",goodbye)
	http.ListenAndServe(":8080",pr)
}
func newPathResolver() *pathResolver{
	return &pathResolver{make(map[string]http.HandlerFunc)}
}
type pathResolver struct{
	handlers map[string]http.HandlerFunc
}
func (p *pathResolver) add(path string,handler http.HandlerFunc){
	p.handlers[path]=handler
}
func (p *pathResolver) ServeHTTP(res http.ResponseWriter,req *http.Request){
	check:=req.Method+" "+req.URL.Path
	for pattern,handlerFunc:=range p.handlers{
		if ok,err:=path.Match(pattern,check);ok&&err==nil{
			handlerFunc(res,req)
			return
		}else if err!=nil{
			fmt.Fprint(res,err)
		}
	}
}
func hello (res http.ResponseWriter,req *http.Request){
	query:=req.URL.Query()
	name:=query.Get("name")
	if name==""{
		name="fenger"
	}
	fmt.Fprintf(res,"hello,my name is ",name)
}
func goodbye(res http.ResponseWriter,req *http.Request){
	path:=req.URL.Path
	parts:=strings.Split(path,"/")
	name :=parts[2]
	if name==""{
		name="fenger"
	}
	fmt.Fprint(res,"goodbye ",name)
}