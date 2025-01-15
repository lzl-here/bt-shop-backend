package verr

type VErr struct {
	error
	code int
	msg  string
}

func New(code int, msg string) *VErr {
	return &VErr{
		code: code,
		msg:  msg,
	}
}

func (v *VErr) Error() string {
	return v.msg
}
func (v *VErr) Code() int {
	return v.code
}
