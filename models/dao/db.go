package dao

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User 用户表
type User struct {
	UID      int    `orm:"auto;pk;column(uid)"`
	Name     string `orm:"column(name)"`
	Password string `orm:"column(password)"`
}

// 数据库连接
func init() {
	aliseName := beego.AppConfig.String("aliseName")
	driverName := beego.AppConfig.String("driverName")
	dataSource := beego.AppConfig.String("dataSource")

	orm.RegisterDataBase(aliseName, driverName, dataSource)
	orm.RegisterModel(new(User))
}

// SelectUserTable 连接数据库
func SelectUserTable(name, passwd string) bool {

	user := new(User)
	o := orm.NewOrm()
	ok := o.QueryTable(user).Filter("name", name).Filter("password", passwd).Exist()

	return ok
}

// Decode 用于将查询出来的
func Decode() {

}
