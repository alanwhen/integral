package models

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
}
