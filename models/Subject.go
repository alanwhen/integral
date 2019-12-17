package models

type Subject struct {
	Id          int
	SubjectName string `orm:"size(50)"`
	ListOrder   int
	Status      byte
}
