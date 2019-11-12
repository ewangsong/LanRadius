package main

import (
	"radiusweb/cmd"
	_ "radiusweb/models"
	_ "radiusweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	jsonConfig := `{
	    "filename" : "/var/log/goradius.log",
	    "maxlines" : 1000,
	    "maxsize"  : 10240
	}`

	logs.SetLogger(logs.AdapterFile, jsonConfig) // 设置日志记录方式：本地文件记录
	logs.EnableFuncCallDepth(true)               // 输出log时能显示输出文件名和行号（非必须）
	beego.BeeLogger.DelLogger("console")
	cmd.Execute()
}

func main() {

	// admin := flag.Bool("admin", false, "Run admin interface")
	// radiusct := flag.Bool("radiusct", false, "Run radius server")
	// help := flag.Bool("help", false, "Print Help")
	// rand.Seed(time.Now().UTC().UnixNano())
	// flag.Parse()

	// if *help == true {
	// 	print_help()
	// 	return
	// }

	// if *admin == true {
	// 	fmt.Println("run beego")
	// 	beego.Run()
	// } else {
	// 	if *radiusct == true {
	// 		fmt.Println("Run radius server....")

	// 		radius.RadiusRun()
	// 	}

	// }

	// print_help()

}
