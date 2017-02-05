package routers

import "github.com/astaxie/beego"
import"chandra_application/controller"


func init() {
	beego.Router("bypass/:id", &controllers.Maincontroller{},"get:Bypass")
	beego.Router("/",&controllers.Maincontroller{})
	beego.Router("/:id", &controllers.Maincontroller{})


}