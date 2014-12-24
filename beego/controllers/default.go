package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	this.Data["Body"] = string(this.Ctx.Input.RequestBody)
	println(this.Ctx.Input.RequestBody, len(this.Ctx.Input.RequestBody))
	this.TplNames = "index.tpl"
}

func (this *MainController) HttpGzip() {
	this.Data["json"] = this.Ctx.Request.Header
	this.ServeJson()
}
