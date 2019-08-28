package controller

type ResponseResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseResource(code int, msg string, data interface{}) (responseResult *ResponseResult) {
	responseResult = &ResponseResult{code, msg, data}
	return
}
