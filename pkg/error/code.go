package error

var (
	BadRequest      = &Error{ErrorCode: "00001", Message: "Fail."}
	ParamBadRequest = &Error{ErrorCode: "00002", Message: "Lack Param Or Param Not Satisfiable."}

	TokenInvalid = &Error{ErrorCode: "10001", Message: "Token Invalid."}
	Unauthorized = &Error{ErrorCode: "10002", Message: "Unauthorized."}

	LoginError = &Error{ErrorCode: "20001", Message: "Password or account wrong."}

	ServerError = &Error{ErrorCode: "30001", Message: "Fail."}
)
