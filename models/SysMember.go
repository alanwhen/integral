package models

import "time"

type SysMember struct {
	Id            int       `orm:"column(member_id)" from:"Id"`
	Username      string    `orm:"column(username)" from:"Username"`
	Password      string    `orm:"column(password)" from:"Password"`
	Email         string    `orm:"column(email)" from:"Email"`
	GroupId       string    `orm:"column(group_id)" from:"GroupId"`
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
}
