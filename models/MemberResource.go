package models

type MemberResource struct {
	Id                    int
	Name                  string          `orm:"size(64)"`
	Parent                *MemberResource `orm:"null;rel(fk)"`
	Rtype                 int
	Seq                   int
	Sons                  []*MemberResource        `orm:"reverse(many)"` //fk 的反向关系
	SonNum                int                      `orm:"-"`
	Icon                  string                   `orm:"size(32)"`
	LinkUrl               string                   `orm:"-"`
	UrlFor                string                   `orm:"size(255)" Json:"-"`
	HtmlDisabled          int                      `orm:"-"`
	Level                 int                      `orm:"-"`
	MemberRoleResourceRel []*MemberRoleResourceRel `orm:"reverse(many)"`
}
