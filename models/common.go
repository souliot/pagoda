package models

type PageQuery struct {
	From  int `orm:"-" json:"-"`
	Limit int `orm:"-" json:"-"`
}

type List struct {
	Total int64       `json:"total"`
	Lists interface{} `json:"lists"`
}
