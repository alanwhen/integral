package models

import "github.com/alanwhen/education-mini/enums"

type JsonResult struct {
	Status enums.JsonResultStatus `json:"status"`
	Tips   string                 `json:"tips"`
	Data   interface{}            `json:"data"`
}

type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int64  `json:"offset"`
	Limit  int    `json:"limit"`
}
