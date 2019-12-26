package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type (
	Teacher struct {
		Id           int
		SchoolId     int
		SubjectId    int
		TeacherName  string `orm:"size()"`
		TeacherNo    string `orm:"size()"`
		Introduction string `orm:"type(text)"`
		Sex          byte
		Avatar       string    `orm:"size(150)"`
		Birthday     time.Time `orm:"type(date)"`
		Mobile       string    `orm:"size(22)"`
		Password     string
		AddTime      int
		Openid       string `orm:"size(50)"`
	}

	TeacherGrid struct {
		Id          int
		SchoolName  string
		SubjectName string
		TeacherName string
		TeacherNo   string
		Sex         byte
		Mobile      string
	}

	TeacherQueryParam struct {
		BaseQueryParam
		TeacherNo string
	}
)

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("db_tb_prefix"), new(Teacher))
}

func (this *Teacher) TableName() string {
	return TeacherTBName()
}

func TeacherTBName() string {
	return "teacher"
}

func TeacherPageList(params *TeacherQueryParam) ([]*TeacherGrid, int64) {
	query := orm.NewOrm()
	data := make([]*TeacherGrid, 0)
	sortOrder := "t.Id"
	switch params.Sort {
	case "Id":
		sortOrder = "t.Id"
	}
	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	sql := fmt.Sprintf(`SELECT sc.school_name,sj.subject_name,t.id,t.teacher_name,t.teacher_no,t.mobile,t.sex
                               FROM %s AS t
                               LEFT JOIN %s sc ON t.school_id = sc.id
                               LEFT JOIN %s sj ON t.subject_id = sj.id
                               WHERE t.teacher_no LIKE '%%%s%%'
                               ORDER BY %s
                               `,
		beego.AppConfig.String("db_dt_prefix")+"teacher",
		beego.AppConfig.String("db_dt_prefix")+"school",
		beego.AppConfig.String("db_dt_prefix")+"subject",
		params.TeacherNo,
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

func TeacherBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(StudentTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func TeacherOne(id int) (*Teacher, error) {
	o := orm.NewOrm()
	m := Teacher{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
