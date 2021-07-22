package models

type ResponseModel struct {
	Status   int         `json:"status"`
	Code     int         `json:"code"`
	Type     string      `json:"type"`
	Message  string      `json:"message"`
	MoreInfo string      `json:"more_info"`
	Data     interface{} `json:"data"`
}
