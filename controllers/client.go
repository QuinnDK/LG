package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type ClientController struct {
	beego.Controller
}

func (this *ClientController) Get() {
	this.TplName = "client.html"
}

func (this *ClientController) Post() {
	options := this.Input().Get("options")
	beego.Info(options)
	//请求检查方法
	if options != "" {
		switch options {
		case "login":
			this.login()
		case "register":
			this.register()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "无对应处理方法！", "time": time.Now().Format("2006-12-12 12:12:12")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "options为空", "time": time.Now().Format("2006-12-12 12:12:12")}
		this.ServeJSON()
		return

	}

}