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

func (a *News) TableName() string {
	return "news"
}

func GetNews(user int) ([]News) {
	o := orm.NewOrm()
	o.Using("default")
	var result []News
	num, err := o.Raw("SELECT * FROM test.news").QueryRows(&result)


	if num < 0 {

	}

	if err == nil {

	}

	return result
}