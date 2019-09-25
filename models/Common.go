package models

import "github.com/alanwhen/education-mini/enums"

type JsonResult struct {
	Status enums.JsonResultStatus `json:"status"`
	Tips   string                 `json:"tips"`
	Data   interface{}            `json:"data"`
}
