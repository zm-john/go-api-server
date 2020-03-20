package handler

// Error 错误结构
type Error struct {
	// 表单字段
	Field string `json:"field"`
	// 错误说明
	Message string `json:"message"`
}

// ErrorResponse 错误时响应结构
// swagger:response error
// in: body
type ErrorResponse struct {
	Message string  `json:"message"`
	Erorrs  []Error `json:"errors,omitempty"`
}

// Response 正常是响应结构
type Response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

// NewErrResponse 新建一个错误响应
func NewErrResponse(message string, errs []Error) *ErrorResponse {
	return &ErrorResponse{message, errs}
}

// NewResponse 新建一个正常响应
func NewResponse(data, meta interface{}) *Response {
	return &Response{data, meta}
}
