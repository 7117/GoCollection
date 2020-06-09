package main

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	// 方式错误
	ErrorRequestNotRecognized = Err{Error: "api not recognized", ErrorCode: "001"}
	// body错误
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	// 内部错误
	ErrorInternalFaults = Err{Error: "internal service error", ErrorCode: "003"}
)
