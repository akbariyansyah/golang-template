package formatter

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(data any) *Response {
	return &Response{Message: "ok", Data: data}
}

func NewErrorResponse(err error) *Response {
	return &Response{Message: err.Error()}
}
