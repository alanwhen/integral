package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Subject struct {
	Id          int
	SubjectName string `orm:"size(50)"`
	ListOrder   int
	Status      byte
}

func init() {
	orm.RegisterModel(new(Subject))
}

type SubjectQueryParam struct {
	BaseQueryParam
	SubjectName string
}

func SubjectTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "subject"
}

func SubjectPageList(params *SubjectQueryParam) ([]*Subject, int64) {
	query := orm.NewOrm().QueryTable(SubjectTBName())
	data := make([]*Subject, 0)

	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	}

	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func SubjectOne(id int) (*Subject, error) {
	o := orm.NewOrm()
	m := Subject{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func SubjectBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(SubjectTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func (this *Subject) TableName() string {
	return SubjectTBName()
}
