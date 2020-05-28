package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,Error:Err{Error:"ErrorRequestBodyParseFailed",ErrorCode:"001"}}
	ErrroNotAuthUser =  ErrorResponse{HttpSC:401,Error:Err{Error:"ErrroNotAuthUser",ErrorCode:"002"}}
)