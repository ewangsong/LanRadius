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
)

func print_help() {
	fmt.Println("使用帮助:")
	fmt.Println("-admin 启动Web UI管理")
	fmt.Println("-radiusct 启动后台radius程序")

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
			//	fmt.Println("Run radius server....")

			radius.RadiusRun()
		}

	}

	print_help()

}
