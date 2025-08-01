package mocking

type LoggingMock struct {
	InfoMessage []string
	InfoError   []string
}

func (l *LoggingMock) MsgInfo(msg string) {
	l.InfoMessage = append(l.InfoMessage, msg)
}

func (l *LoggingMock) ErrInfo(err error, str string) {
	l.InfoError = append(l.InfoError, str+" : "+err.Error())
}
