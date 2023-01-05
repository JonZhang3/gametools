package app

type BizError struct {
	Message string
}

func NewBizError(message string) *BizError {
	return &BizError{message}
}

func (e *BizError) Error() string {
	return e.Message
}
