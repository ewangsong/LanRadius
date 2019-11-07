package main

import (
	"flag"
	"fmt"
	"math/rand"
	_ "radiusweb/models"
	"radiusweb/radius"
	_ "radiusweb/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func print_help() {
	fmt.Println("使用帮助:")
	fmt.Println("-admin 启动Web UI管理")
	fmt.Println("-radiusct 启动后台radius程序")

}
func init() {
	jsonConfig := `{
	    "filename" : "/var/log/goradius.log",
	    "maxlines" : 1000,
	    "maxsize"  : 10240
	}`

	logs.SetLogger(logs.AdapterFile, jsonConfig) // 设置日志记录方式：本地文件记录
	logs.EnableFuncCallDepth(true)               // 输出log时能显示输出文件名和行号（非必须）
	beego.BeeLogger.DelLogger("console")
}

func main() {

	admin := flag.Bool("admin", false, "Run admin interface")
	radiusct := flag.Bool("radiusct", false, "Run radius server")
	help := flag.Bool("help", false, "Print Help")
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()

	if *help == true {
		print_help()
		return
	}

	if *admin == true {
		fmt.Println("run beego")
		beego.Run()
	} else {
		if *radiusct == true {
			fmt.Println("Run radius server....")

			radius.RadiusRun()
		}

	}

	print_help()

}
