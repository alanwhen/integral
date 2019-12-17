package models

import "time"

type Student struct {
	Id           int    `orm:"column(student_id)" from:"Id"`
	StudentNo    string `orm:"size(20)"`
	StudentName  string `orm:"size(20)"`
	Avatar       string `orm:"size(150)"`
	User         *User  `orm:"rel(fk)"`
	School       *School
	Grade        byte
	Sex          byte
	Birthday     time.Time `orm:"type(date)"`
	IfDel        byte
	ProvinceId   int
	CityId       int
	CountyId     int
	TownId       int
	Address      string `orm:"size(150)"`
	Lat          string `orm:"size(50)"`
	Lng          string `orm:"size(50)"`
	Relation     int
	AddTime      int
	ParentName   string `orm:"size(50)"`
	ParentMobile string `orm:"size(50)"`
}
