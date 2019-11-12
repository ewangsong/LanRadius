package controllers

import (
	"fmt"
	"os/exec"
	"radiusweb/libs"
	"radiusweb/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//插入
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "123456"
	// user.Passwd = "wangsong"
	// _, err := o.Insert(&user)
	// if err != nil {
	// 	beego.Info("插入错误", err)
	// 	return
	// }
	//	查询
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "wangsong"
	// err := o.Read(&user, "Name")
	// if err != nil {
	// 	//	beego.Info("查询失败", err)
	// 	return
	// }
	// beego.Info(user.Passwd)
	// p := user.Passwd
	//删除
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "123456"
	// _, err := o.Delete(&user, "Name")
	// if err != nil {
	// 	beego.Info("worng", err)
	// 	return
	// }

	//更新
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "MyName"
	// if o.Read(&user, "Name") == nil {
	// 	user.Passwd = "55555"
	// 	_, err := o.Update(&user)
	// 	if err != nil {
	// 		beego.Info(err)
	// 	}
	// }
	//注册
	c.Redirect("/login.html", 302)
}
func (c *MainController) Post() {
	//	c.Data["data"] = "一一一一"
	//	c.TplName = "index.html"
}

//登入页面方法
func (c *MainController) Login() {
	c.TplName = "login.html"
}

func (c *MainController) Getadmin() {
	c.TplName = "login.html"
	Name1 := c.GetString("username")
	Passwd1 := c.GetString("password")
	o := orm.NewOrm()
	admin := models.WsAdmin{}
	admin.Name = Name1
	err := o.Read(&admin, "Name")
	beego.Info(Passwd1, admin.Password)
	if err != nil || admin.Password != Passwd1 {
		c.Ctx.WriteString("账号有误，请检查账号")
		return
	}

	c.Redirect("/admin/dashboard", 302)
}

//主页方法
func (c *MainController) Getindex() {
	c.Layout = "layout_base.html"
	c.TplName = "template_ui.html"

}

//关于about
func (c *MainController) ShowAbout() {

	c.Layout = "layout_base.html"
	// if c.LayoutSections == nil {
	// 	c.LayoutSections = make(map[string]string)
	// }
	// c.LayoutSections["ui"] = "template_ui.html"
	c.TplName = "about.html"
}

//修改密码
func (c *MainController) ShowChangePassword() {
	c.Layout = "layout_base.html"

	c.TplName = "password.html"
}

func (c *MainController) PostChangePassword() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.LayoutSections["HeadCss"] = "password_head.html"
	c.TplName = "password.html"
	p1 := c.GetString("tr_user_pass")
	p2 := c.GetString("tr_user_pass_chk")
	if models.AdminUpdate(p1, p2) == "ok" {
		c.Data["T"] = "更改成功"
	} else {
		c.Data["T"] = "更改失败，请重新更改"
	}

}

//退出登入
func (c *MainController) LogOut() {
	c.Redirect("/login", 302)

}

//操作日志界面
func (c *MainController) ShowLog() { //get请求
	c.Layout = "layout_base.html"
	c.TplName = "log.html"
	pno, _ := c.GetInt("pno") //获取当前请求页
	var ShowLog []models.WsLog
	var conditions []string = make([]string, 2) //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions                   //定义一个分页对象
	po.TableName = "ws_log"                     //指定分页的表名
	po.EnableFirstLastLink = true               //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                  //是否显示上一页下一页 默认为false
	po.Conditions = conditions                  // 传递分页条件 默认全表
	po.Currentpage = int(pno)                   //传递当前页数,默认为1
	po.PageSize = 10                            //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&ShowLog)      //把当前页面的数据序列化进一个切片内
	c.Data["ShowLog"] = ShowLog //把当前页面的数据传递到前台
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem

}

func (c *MainController) PostLog() { //post请求
	c.Layout = "layout_base.html"
	c.TplName = "log.html"
}

//系统服务页

func (c *MainController) ShowSuperrpc() {
	c.Layout = "layout_base.html"
	c.TplName = "wait.html"
}

//控制面板
func (c *MainController) ShowDashboard() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["CPU"] = c.GetCpu()
	c.Data["ServerTime"] = c.GetServerTime()
	c.Data["MemInfo"] = c.GetMemInfo()
	c.Data["Disk"] = c.GetDiskUsage()

	c.TplName = "dashboard.html"
}
func (c *MainController) GetServerTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func (c *MainController) GetCliOutput(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}

func (c *MainController) GetMemInfo() [][]string {
	cmd := "cat /proc/meminfo"
	tmp := c.GetCliOutput(cmd)
	var ret [][]string
	arr := strings.Split(tmp, "\n")
	for _, v := range arr {
		if strings.HasPrefix(v, "MemTotal") || strings.HasPrefix(v, "MemFree") || strings.HasPrefix(v, "Cached") {
			//v = strings.Replace(v,"  ","",-1)
			//fmt.Println(v)
			fields := strings.Fields(v)
			fmt.Println(fields)
			a := []string{fields[0], fields[1]}
			ret = append(ret, a)
		}
	}

	return ret
}

func (c *MainController) GetCpu() string {

	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	return c.GetCliOutput(cmd)
}

func (c *MainController) GetDiskUsage() *libs.DiskStatus {
	dk := libs.DiskUsage("/")
	return &dk
}
