package routers

import (
	"github.com/astaxie/beego"
	"go_code/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gzip", &controllers.MainController{}, "get:HttpGzip")
}
