package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type School struct {
	Id            int
	SchoolName    string `orm:"size(50)"`
	SchoolNo      string `orm:"size(20)"`
	ProvinceId    int
	CityId        int
	CountyId      int
	TownId        int
	Lat           string `orm:"size(50)"`
	Lng           string `orm:"size(50)"`
	Introduction  string `orm:"type(text)"`
	AddTime       int
	ListOrder     int
	Address       string `orm:"size(50)"`
	AddressDetail string `orm:"size(200)"`
}

func init() {
	orm.RegisterModel(new(School))
}

type SchoolQueryParam struct {
	BaseQueryParam
}

func SchoolTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "school"
}

func SchoolPageList(params *SchoolQueryParam) ([]*School, int64) {
	query := orm.NewOrm().QueryTable(SchoolTBName())
	data := make([]*School, 0)

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

func SchoolDataList(params *SchoolQueryParam) []*School {
	params.Limit = -1
	data, _ := SchoolPageList(params)
	return data
}

func SchoolOne(id int) (*School, error) {
	o := orm.NewOrm()
	m := School{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func SchoolBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(SchoolTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func (this *School) TableName() string {
	return SchoolTBName()
}
