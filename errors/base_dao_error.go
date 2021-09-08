package errors

import "fmt"

type BizError struct {
	Code    int
	Message string
}

func NewBizError(msg string) *BizError {
	return &BizError{
		Code:    501,
		Message: msg,
	}
}

func NewDaoError(msg string) *BizError {
	return &BizError{
		502,
		msg,
	}
}

func (err *BizError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", err.Code, err.Message)
}
