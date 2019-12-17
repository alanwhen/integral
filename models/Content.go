package models

type Content struct {
	Id              int
	ArticleCategory *ArticleCategory `orm:"rel(fk)"`
	Type            int8
	Title           string
	Cover           string
	Content         string
	Url             string
	Status          byte
	Member          *SysMember
	ListOrder       int
	AddTime         int
}
