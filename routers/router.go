package routers

import (
	"sitepointgoapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/news", &controllers.MainController{}, "get:NewsListAction")
	beego.Router("/news/:id([0-9]+)", &controllers.MainController{}, "get:NewItemAction")
}
