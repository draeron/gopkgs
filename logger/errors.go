package logger

import "go.uber.org/zap"

func (l *Logger) LogIfErr(err error) {
	if err != nil {
		l.WithOptions(zap.AddCallerSkip(1)).With(zap.Error(err)).Error("error occured")
	}
}

func (l *Logger) Sugar() *SugaredLogger {
	return &SugaredLogger{l.Logger.Sugar()}
}

func (l *Logger) StopIfErr(err error) {
	if err != nil {
		l.WithOptions(zap.AddCallerSkip(1)).With(zap.Error(err)).Fatal("error occured")
	}
}

func (l *SugaredLogger) LogIfErr(err error) {
	if err != nil {
		l.Desugar().WithOptions(zap.AddCallerSkip(1)).With(zap.Error(err)).Error("error occured")
	}
}

func (l *SugaredLogger) StopIfErr(err error) {
	if err != nil {
		l.Desugar().WithOptions(zap.AddCallerSkip(1)).With(zap.Error(err)).Fatal("error occured")
	}
}
