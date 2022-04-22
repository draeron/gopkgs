package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	*zap.Logger
}

type SugaredLogger struct {
	*zap.SugaredLogger
}

func NewZap(name string) *SugaredLogger {
	// for !flag.Parsed() {
	//	<- time.After(time.Millisecond * 50)
	// }

	l, err := zap.NewDevelopment()
	l = l.WithOptions(zap.AddStacktrace(zap.FatalLevel))

	if err != nil {
		zap.S().Panic(err)
	}

	return &SugaredLogger{l.Named(name).Sugar()}
}

func (l *SugaredLogger) ErrorIf(err error, fmt string, args ...interface{}) {
	if err != nil {
		l.Errorf(fmt, args...)
	}
}
