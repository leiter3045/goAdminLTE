package traits

type ErrorReport struct {
	_errorMsg string
}

func (u ErrorReport) GetError() (string) {
	return u._errorMsg
}

func (u *ErrorReport) SetError(msg string) {
	u._errorMsg = msg
}