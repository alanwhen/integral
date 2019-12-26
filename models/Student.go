package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type (
	Student struct {
		Id           int     `orm:"column(student_id)" from:"Id"`
		StudentNo    string  `orm:"size(20),column(student_no)" from:"StudentNo"`
		StudentName  string  `orm:"size(20)"`
		Avatar       string  `orm:"size(150)"`
		User         *User   `orm:"rel(fk)"`
		School       *School `orm:"rel(fk)"`
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

	StudentGrid struct {
		Id           int
		StudentNo    string
		StudentName  string
		Sex          byte
		Grade        byte
		SchoolName   string
		ParentName   string
		ParentMobile string
	}

	StudentQueryParam struct {
		BaseQueryParam
		StudentNo string
	}
)

func init() {
	orm.RegisterModel(new(Student))
}

func StudentTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "student"
}

func StudentPageList(params *StudentQueryParam) ([]*StudentGrid, int64) {
	query := orm.NewOrm()

	data := make([]*StudentGrid, 0)

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

	sql := fmt.Sprintf(`SELECT st.student_id as id,st.student_no,st.sex,st.grade,st.student_name,st.parent_name,st.parent_mobile,sc.school_name
                               FROM %s AS st
                               LEFT JOIN %s as sc ON st.school_id = sc.id
                               WHERE st.student_no LIKE '%%%s%%'
                               ORDER BY %s
                              `,
		beego.AppConfig.String("db_dt_prefix")+"student",
		beego.AppConfig.String("db_dt_prefix")+"school",
		params.StudentNo,
		sortOrder,
	)

	total, err := query.Raw(sql).QueryRows(&data)
	if err != nil {
		return nil, 0
	}

	sql = sql + fmt.Sprintf(" LIMIT %d, %d", params.Offset, params.Limit)
	_, err = query.Raw(sql).QueryRows(&data)
	if err != nil {
		return nil, 0
	}
	return data, total
}

func StudentDataList(params *StudentQueryParam) []*StudentGrid {
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
