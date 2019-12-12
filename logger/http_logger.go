package logger

type Http struct {
	*SugaredLogger
}

func (l *Http) Print(v ...interface{}) {
	l.Info(v...)
}
