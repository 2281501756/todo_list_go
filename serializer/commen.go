package serializer

// 返回值序列化器

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Err    string      `json:"err"`
}

type ResponseError struct {
	Msg interface{} `json:"msg"`
}
