package models

import "github.com/astaxie/beego/orm"

type SiteConfig struct {
	ConfigKey   string `orm:"column(config_key)" from:"ConfigKey"`
	ConfigValue string `orm:"column(config_value)" from:"ConfigValue"`
	ConfigDesc  string `orm:"column(config_desc)" from:"ConfigDesc"`
	Required    int    `orm:"column(required)" from:"Required"`
	Text        string `orm:"column(text)" from:"Text"`
	CharType    string `orm:"column(char_type)" from:"CharType"`
}

func init() {
	orm.RegisterModel(new(SiteConfig))
}

func SiteConfigTBName() string {
	return "site_config"
}

func (this *SiteConfig) TableName() string {
	return SiteConfigTBName()
}

func SiteConfigList() []*SiteConfig {
	query := orm.NewOrm().QueryTable(SiteConfigTBName())
	data := make([]*SiteConfig, 0)
	query.All(&data)
	return data
}
