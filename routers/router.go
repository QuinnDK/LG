package routers

import (
	"LG/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/client", &controllers.ClientController{})
}
