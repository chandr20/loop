package main

import (

        _ "chandra_application/routers"
	//"github.com/astaxie/beego/orm"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)


func init(){
	orm.RegisterDataBase("default", "mysql", "root:devstack@tcp(127.0.0.1:3306)/chan_app")
	name := "default"

	 //Drop table and re-create.
	 force := false

	//// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)


	}



}

func main() {

	beego.Run()

}
