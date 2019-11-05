package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func BasInsert(name, ip_add, secret, port string) bool {
	bas := WsBas{}

	o := orm.NewOrm()
	bas.Name = name
	bas.IpAddr = ip_add
	bas.Secret = secret
	bas.Port = port
	_, err := o.Insert(&bas)
	if err != nil {
		beego.Info("bas插入错误", err)
		return false
	} else {
		return true
	}

}

//添加用户数据库操作

func UserInsert(realname, name, password string) {
	user := WsUsers{RealName: realname, Name: name, Password: password}
	o := orm.NewOrm()
	beego.Info(user)
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("用户插入错误", err)
	}

}

//插入操作日志
func LogInsert() {

}
