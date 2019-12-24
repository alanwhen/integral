package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id                   int
	CourseName           string
	CourseNo             string
	School               *School `orm:"rel(fk)"`
	IsRecommend          byte
	Teacher              *Teacher `orm:"rel(fk)"`
	Subject              *Subject `orm:"rel(fk)"`
	Professor            string
	Place                string
	WeekTime             byte
	IfWeekTime           byte
	Type                 byte
	CourseDuration       int
	Lesson               int16
	IfShowLesson         byte
	Logo                 string
	Banner               string
	Phone                string `orm:"size(22)"`
	AgeRange             int16
	GradeRange           int16
	Price                float64 `orm:"digits(10);decimals(2)"`
	MarketPrice          float64 `orm:"digits(10);decimals(2)"`
	PriceDesc            string
	MarketPriceDesc      string
	Plan                 string
	Description          string
	Effective            string
	Aims                 string
	EnrollableNum        int16
	IfShowEnroll         byte
	Enrolment            int16
	AddTime              int
	RegistrationTime     int
	RegistrationDeadline int
	StartTime            int
	EndTime              int
	Status               byte
	ListOrder            int
}

type CourseQueryParam struct {
	BaseQueryParam
}

func init() {
	orm.RegisterModel(new(Course))
}

func CourseTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "course"
}

func CoursePageList(params *CourseQueryParam) ([]*Course, int64) {
	query := orm.NewOrm().QueryTable(CourseTBName())
	data := make([]*Course, 0)

	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "CourseNo":
		sortOrder = "CourseNo"
	}

	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

func CourseDataList(params *CourseQueryParam) []*Course {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := CoursePageList(params)
	return data
}

func CourseOne(id int) (*Course, error) {
	o := orm.NewOrm()
	m := Course{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
