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
	fmt.Println("Useage:")
	fmt.Println("softradius -admin")
	fmt.Println("softradius -radacct")

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
