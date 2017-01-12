package controllers

import (
	"github.com/astaxie/beego"
	"sitepointgoapp/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) NewItemAction() {
	c.Data["Title"] = "Hello page"

	number, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	o := orm.NewOrm()

	user := models.Users{Id: number}
	err := o.Read(&user)

	if(err != nil){
		fmt.Println("News not found")
	}

	beego.Debug(user)

	c.Data["Text"] = user.Info

	c.TplName = "default/hello-sitepoint.tpl"
}

func (main *MainController) TestAction() {
	main.Data["Title"] = "Test"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = "One"

	o := orm.NewOrm()

	user := new(models.Users)
	user.Name = "string"
	user.Info = "rand"

	fmt.Println(o.Insert(user))


	//err := models.AddUser("test")
	//if(err == nil){
	//	main.Data["Result"] = true
	//}

	main.TplName = "default/hello-sitepoint.tpl"
}

func (main *MainController) NewsListAction() {
	main.Data["Title"] = "Test"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = "One"

	o := orm.NewOrm()

	user := new(models.Users)
	user.Name = "string"
	user.Info = "rand"

	fmt.Println(o.Insert(user))


	//err := models.AddUser("test")
	//if(err == nil){
	//	main.Data["Result"] = true
	//}

	main.TplName = "default/hello-sitepoint.tpl"
}
