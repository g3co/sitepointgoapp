package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id int
	Name string
	Info string
}

type News struct {
	Id int
	Name string
	Prev string
	Content string
}


func init() {
	// Need to register model in init
	orm.RegisterModel(new(Users), new(News))
}

func AddUser(username string) (error) {
	o := orm.NewOrm()
	o.Using("default")

	err := o.Raw("INSERT INTO `first` (`name`, `info`) VALUES (?, 'qwrq');", username).QueryRow()

	return err
}