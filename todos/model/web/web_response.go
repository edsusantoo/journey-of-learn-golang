package web

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type MessageResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
	Token  any    `json:"token"`
}
