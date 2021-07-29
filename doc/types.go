package doc

type ApiResponse struct {
	Status   int         `json:"status,omitempty"`
	Code     int         `json:"code,omitempty"`
	Type     string      `json:"type,omitempty"`
	Message  string      `json:"message,omitempty"`
	MoreInfo string      `json:"more_info,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
