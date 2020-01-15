package error

var (
	TokenInvalid = &Error{ErrorCode:"10000001",Message:"Token Invalid."}

	LoginError = &Error{ErrorCode:"00000001", Message:""}
)
