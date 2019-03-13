package utils

const (
	ErrSuccess         = 0
	ErrAccessDenied    = 1
	ErrNotFound        = 2
	ErrExist           = 3
	ErrInvalidArgument = 4
	ErrInternal        = 5
	ErrDBOperation     = 6
	ErrFileOperation   = 7
	ErrNetwork         = 8
	ErrMaxLimited      = 9
	ErrChallenge       = 10
)

var (
	ErrCodeMsg map[int]string = map[int]string{
		ErrSuccess:         "success",
		ErrAccessDenied:    "access denied",
		ErrNotFound:        "not found",
		ErrExist:           "exist",
		ErrInvalidArgument: "invalid argument",
		ErrInternal:        "general error",
		ErrDBOperation:     "database error",
		ErrFileOperation:   "file error",
		ErrNetwork:         "network error",
		ErrMaxLimited:      "reach the max limitation",
		ErrChallenge:       "need challenge",
	}
)

// Result represents HTTP response body.
type Result struct {
	Code int         `json:"error"`    // return code, 0 for succ
	Msg  string      `json:"errormsg"` // message
	Data interface{} `json:"data"`     // data object
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
}
