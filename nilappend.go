package main
import("log")
func main() {
	res :=[]*Cyeam{}
	a:=&Cyeam{Owner:"A"}
res=append(res,a)

res=append(res,nil)

res=append(res,a)
log.Println(res)
}

type Cyeam struct {
	Owner string
}