package log


type Logging interface {
	MsgInfo(msg string)
	ErrInfo(err error, str string)
}