package main

import (
	_ "sitepointgoapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:01490848@tcp(localhost:3307)/test?charset=utf8")
}

func main() {
	beego.Run()
}