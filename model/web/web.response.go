package web

type WebResponse struct {
	Code     int         `json:"code"`
	Status   string      `json:"status,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    string      `json:"error,omitempty"`
	Messsage string      `json:"messsage,omitempty"`
}
