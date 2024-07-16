package model

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Responnses []Response

type MessageResponse struct {
	Data     interface{} `json:"data"`
	Errors   Responnses  `json:errors"`
	Messages Responnses  `json:messages`
}
