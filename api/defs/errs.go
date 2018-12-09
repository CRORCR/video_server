package defs

/**
 * @desc    返回错误码
 * @author Ipencil
 * @create 2018/12/9
 */
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct{
	HttpSC int `json:"sc"`
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,Error:Err{Error:"Request body is not correct.",ErrorCode:"001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC:401,Error:Err{Error:"User authentication failed.",ErrorCode:"002"}}
)

