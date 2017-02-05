package models

import (
	"github.com/astaxie/beego/orm"

	//t6"strconv"
)

type Student struct {

	Id         int    `orm:"column(id);auto;null"`
	School     string `orm:"column(school);size(25);null"`
	Course     string `orm:"column(course);size(25);null"`
	Gender     string `orm:"column(gender);size(25);null"`


}

func init(){
	orm.RegisterModel(new(Student))

}
func (*Student)AddData(m *Student)(id int64,err error){
	o:=orm.NewOrm()
	id, err = o.Insert(m)
	return

}

func (*Student)GetData(id int)(v *Student,err error){
	o:=orm.NewOrm()
	v = &Student{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


