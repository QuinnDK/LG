package controllers

import (
	"LG/models"
	"github.com/astaxie/beego"
	"time"
)

//登录
func (this *ClientController) login() {

	//从前端获取数据信息
	pwd := this.Input().Get("Pwd")
	tel := this.Input().Get("Tel")

	beego.Info(pwd)
	beego.Info(tel)

	if tel == "" || pwd == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不可为空", "time": time.Now().Format("2006-12-12 12:12:12")}
		this.ServeJSON()
		return
	}

	//判断该手机号是否注册
	user, err := models.GetUserByTel(tel)

	if err != nil { //如果有错误
		beego.Info(err)
	}
	if user != nil { //如果有该用户
		if user.Pwd == pwd {
			user.Pwd = ""
			this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空，请重新登录", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//注册
func (this *ClientController) register() {
	tel := this.Input().Get("Tel")
	beego.Info(tel)
	if tel == "" { //如果手机号为空
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//判断该手机号是否已经注册
	user, err := models.GetUserByTel(tel)
	if err != nil { //如果有错误
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if user != nil { //如果有该用户
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "该手机号已注册！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}
