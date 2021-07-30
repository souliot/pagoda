package models

type PageQuery struct {
	Page     int `orm:"-" json:"page,omitempty"`
	PageSize int `orm:"-" json:"pageSize,omitempty"`
}

type List struct {
	Total int64       `json:"total"`
	Lists interface{} `json:"lists"`
}
