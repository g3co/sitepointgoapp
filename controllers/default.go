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
	number, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	newItem := models.News{Id: number}
	o.Read(&newItem)

	c.Data["Title"] = newItem.Name
	c.Data["Text"] = newItem.Content

	c.TplName = "default/news_item.tpl"

}

func (main *MainController) NewsListAction() {
	main.Data["Title"] = "Test"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = "One"

	news := models.GetNews(1)

	main.Data["NewsList"] = news

	main.TplName = "default/news_list.tpl"
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

	main.TplName = "default/hello-sitepoint.tpl"
}