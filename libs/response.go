package libs

type Response struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`
}

func ApiResource(code int64, objects interface{}, msg string) (r *Response) {
	r = &Response{Code: code, Data: objects, Msg: msg}
	return
}
