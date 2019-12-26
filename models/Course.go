package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type (
	Course struct {
		Id                   int
		CourseName           string
		CourseNo             string
		SchoolId             int
		IsRecommend          byte
		TeacherId            int
		SubjectId            int
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

	CourseGrid struct {
		Id          int
		CourseNo    string
		CourseName  string
		SchoolName  string
		SubjectName string
		Lesson      int
		Price       float64
		StartTime   int
		EndTime     int
		ListOrder   int
	}

	CourseQueryParam struct {
		BaseQueryParam
		CourseNo string
	}
)

func init() {
	orm.RegisterModel(new(Course))
}

func CourseTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "course"
}

func CoursePageList(params *CourseQueryParam) ([]*CourseGrid, int64) {
	o := orm.NewOrm()
	data := make([]*CourseGrid, 0)

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

	sql := fmt.Sprintf(`SELECT cs.*,sc.school_name,sj.subject_name
                               FROM %s AS cs
                               LEFT JOIN %s as sc ON cs.school_id = sc.id
                               LEFT JOIN %s as sj ON cs.subject_id = sj.id
                               WHERE cs.course_no LIKE '%%%s%%'
                               ORDER BY %s
                              `,
		beego.AppConfig.String("db_dt_prefix")+"course",
		beego.AppConfig.String("db_dt_prefix")+"school",
		beego.AppConfig.String("db_dt_prefix")+"subject",
		params.CourseNo,
		sortOrder,
	)

	total, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		return nil, 0
	}

	sql = sql + fmt.Sprintf(" LIMIT %d, %d", params.Offset, params.Limit)
	_, err = o.Raw(sql).QueryRows(&data)
	if err != nil {
		return nil, 0
	}
	return data, total
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

func CourseBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(CourseTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func (this *Course) TableName() string {
	return CourseTBName()
}
