package controllers

import (
	"github.com/astaxie/beego"
	"chandra_application/models"
	//"fmt"
        "chandra_application/pkgs/client"
	"encoding/json"
	"strconv"
)

type Maincontroller struct{
	beego.Controller
}

func (this *Maincontroller)Post(){
	var student models.Student
	input:= this.Ctx.Input.RequestBody
	err:=json.Unmarshal(input,&student)
	if err!=nil{
		beego.Info(err)
	}else {

		if _,err := student.AddData(&student); err == nil {
			this.Ctx.Output.SetStatus(200)
			this.Data["json"] = student
		} else {
			this.Data["json"] = err
		}
	}

this.ServeJSON()
}

func (this *Maincontroller)Get(){
       idstr:=this.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idstr)
	var  student models.Student
	data,err := student.GetData(id)
	if err!=nil{
		beego.Info(err)
	}else {
		this.Data["json"]=data
	}
	this.ServeJSON()


}

func (this *Maincontroller)Bypass() {

	var cliententry client.Comm

	cliententry.Id = this.Ctx.Input.Param(":id")

	cliententry.Method = "GET"

	cliententry.Url = beego.AppConfig.String("buildcontroller")

	res,err := cliententry.Conn(&cliententry)
	if err!=nil {
		beego.Info(err)
	}else{
		var super models.Student

		json.Unmarshal(res,&super)
		this.Data["json"] = super



}
this.ServeJSON()

}