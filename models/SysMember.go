package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysMember struct {
	Id            int       `orm:"column(member_id)" from:"Id"`
	Username      string    `orm:"column(username)" from:"Username"`
	Password      string    `orm:"column(password)" from:"Password"`
	Email         string    `orm:"column(email)" from:"Email"`
	GroupId       int       `orm:"column(group_id)" from:"GroupId"`
	OpenId        string    `orm:"column(open_id)" from:"OpenId"`
	Avatar        string    `orm:"column(avatar)" from:"Avatar"`
	RegIp         string    `orm:"column(reg_ip)" from:"RegIp"`
	RegTime       time.Time `orm:"column(reg_time)" from:"RegTime"`
	LastLoginIp   string    `orm:"column(last_login_ip)" from:"LastLoginIp"`
	LastLoginTime time.Time `orm:"column(last_login_time)" from:"LastLoginTime"`
	Encrypt       string    `orm:"column(encrypt)" from:"Encrypt"`
	IfLock        int       `orm:"column(if_lock)" from:"IfLock"`
	FullName      string    `orm:"column(fullname)" from:"FullName"`
	Mobile        string    `orm:"column(mobile)" from:"Mobile"`
	Birth         int       `orm:"column(birth)" from:"Birth"`
	Sign          string    `orm:"column(sign)" from:"Sign"`
}

type SysMemberQueryParam struct {
	BaseQueryParam
	UserName string
	Mobile   string
	Email    string
}

func init() {
	orm.RegisterModel(new(SysMember))
}

func SysMemberTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "sys_member"
}

func (this *SysMember) TableName() string {
	return SysMemberTBName()
}

func SysMemberOne(id int) (*SysMember, error) {
	o := orm.NewOrm()
	m := SysMember{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func SysMemberOneByUserName(username string) (*SysMember, error) {
	m := SysMember{}
	err := orm.NewOrm().QueryTable(SysMemberTBName()).Filter("username", username).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func SysMemberPageList(params *SysMemberQueryParam) {
	//query := orm.NewOrm().QueryTable(SysMemberTBName())
	//data := make([]*SysMember, 0)
	//
	//sortOrder := "Id"
	//switch params.Sort {
	//case "Id":
	//	sortOrder = "Id"
	//}
	//
	//if params.Order == "desc" {
	//	sortOrder = "-" + sortOrder
	//}
	//
	//query = query.Filter("Username__istartswith", params.UserName)
	//query = query.Filter("Mobile__istartswith", params.Mobile)
	//query = query.Filter("Email_istartswith", params.Email)
	//
	//total, _ := query.Count()
	//query.OrderBy(sortOrder).Limit()
}
