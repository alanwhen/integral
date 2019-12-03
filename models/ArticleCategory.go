package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleCategory struct {
	Id        int              `orm:"column(cate_id)"`
	Parent    *ArticleCategory `orm:"null;rel(fk)"`
	Name      string           `orm:"column(cate_name)" from:"Name"`
	ListOrder int              `orm:"column(list_order)" from:"ListOrder"`
	Type      int
	Sons      []*ArticleCategory `orm:"reverse(many)"`
	SonNum    int                `orm:"-"`
	IsSystem  int                `orm:"column(is_system)" from:"IsSystem"`
	Level     int                `orm:"-"`
}

func init() {
	orm.RegisterModel(new(ArticleCategory))
}

func (this *ArticleCategory) TableName() string {
	return ArticleCategoryTBName()
}

func ArticleCategoryTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "article_category"
}

func ArticleCategoryOne(id int) (*ArticleCategory, error) {
	o := orm.NewOrm()
	m := ArticleCategory{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func ArticleCategoryTreeGrid() []*ArticleCategory {
	o := orm.NewOrm()
	query := o.QueryTable(ArticleCategoryTBName()).OrderBy("list_order", "cate_id")
	list := make([]*ArticleCategory, 0)
	query.All(&list)
	return articleCategoryList2TreeGrid(list)
}

func articleCategoryList2TreeGrid(list []*ArticleCategory) []*ArticleCategory {
	result := make([]*ArticleCategory, 0)
	for _, item := range list {
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = articleCategoryAddSons(item, list, result)
		}
	}
}

func articleCategoryAddSons(cur *ArticleCategory, list, result []*ArticleCategory) []*ArticleCategory {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = articleCategoryAddSons(item, list, result)
		}
	}
	return result
}
