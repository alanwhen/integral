package models

import "github.com/alanwhen/education-mini/enums"

type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"tips"`
	Obj  interface{}          `json:"obj"`
}

type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int64  `json:"offset"`
	Limit  int    `json:"limit"`
}
