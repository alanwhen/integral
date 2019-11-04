package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserQueryParam struct {
	BaseQueryParam
}

type User struct {
	Id       int64  `orm:"column(user_id)" from:"Id"`
	Mobile   string `orm:"column(mobile)" from:"Mobile"`
	RealName string `orm:"column(realname)" from:"RealName"`
	Avatar   string `orm:"column(avatar)" from:"Avatar"`
	Sex      int    `orm:"column(sex)" from:"Sex"`
	Password string `orm:"column(password)" from:"Password"`
	AddTime  int    `orm:"column(add_time)" from:"AddTime"`
	Lat      string `orm:"column(lat)" from:"Lat"`
	Lng      string `orm:"column(lng)" from:"Lng"`
}

func init() {
	orm.RegisterModel(new(User))
}

func UserTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "user"
}

func (this *User) TableName() string {
	return UserTBName()
}

func UserPageList(params *UserQueryParam) ([]*User, int64) {
	query := orm.NewOrm().QueryTable(UserTBName())
	data := make([]*User, 0)
	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "Mobile":
		sortOrder = "Mobile"
	}

	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}
	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func UserPageDataList(params *UserQueryParam) []*User {
	params.Limit = -1
	params.Order = "desc"
	data, _ := UserPageList(params)
	return data
}

func UserOne(id int64) (*User, error) {
	o := orm.NewOrm()
	m := User{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func UserOneByMobile(mobile string) (*User, error) {
	m := User{}
	err := orm.NewOrm().QueryTable(UserTBName()).Filter("mobile", mobile).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
