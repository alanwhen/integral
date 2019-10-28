package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoginTrace struct {
	Id         int64  `orm:"column(id)" form:"Id"`
	Username   string `orm:"column(username)" from:"Username"`
	RemoteAddr string `orm:"column(ip)" from:"RemoteAddr"`
	LoginTime  int64  `orm:"column(add_time)" from:"LoginTime"`
}

type LoginTraceQueryParam struct {
	BaseQueryParam
	Username string
}

func init() {
	orm.RegisterModel(new(LoginTrace))
}

func LoginTraceTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "admin_login_logs"
}

func (this *LoginTrace) TableName() string {
	return LoginTraceTBName()
}

func LoginTracePageList(params *LoginTraceQueryParam) ([]*LoginTrace, int64) {
	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	}
	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	query := orm.NewOrm().QueryTable(LoginTraceTBName())
	data := make([]*LoginTrace, 0)
	query = query.Filter("Username__istartswith", params.Username)

	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func LoginTraceDataList(params *LoginTraceQueryParam) []*LoginTrace {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := LoginTracePageList(params)
	return data
}

func LoginTraceAdd(_user string, _remote_add string, _login_time time.Time) error {
	m := LoginTrace{Username: _user, RemoteAddr: _remote_add, LoginTime: _login_time.Unix()}

	o := orm.NewOrm()
	if _, err := o.Insert(&m); err == nil {
		return nil
	} else {
		return err
	}
}
