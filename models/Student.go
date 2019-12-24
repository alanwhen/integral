package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

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

type StudentQueryParam struct {
	BaseQueryParam
	StudentNo string
}

func init() {
	orm.RegisterModel(new(Student))
}

func StudentTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "student"
}

func StudentPageList(params *StudentQueryParam) ([]*Student, int64) {
	query := orm.NewOrm().QueryTable(StudentTBName())
	data := make([]*Student, 0)

	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "StudentNo":
		sortOrder = "student_no"
	}
	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	query = query.Filter("StudentNO__istartswitch", params.StudentNo)

	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func StudentDataList(params *StudentQueryParam) []*Student {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := StudentPageList(params)
	return data
}

func StudentBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(StudentTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func StudentOne(id int) (*Student, error) {
	o := orm.NewOrm()
	m := Student{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (this *Student) TableName() string {
	return StudentTBName()
}
