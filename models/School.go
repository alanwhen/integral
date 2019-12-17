package models

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
