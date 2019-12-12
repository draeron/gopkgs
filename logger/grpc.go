package logger

/*
  GRPC log wrapper
*/

type Grpc struct {
	*SugaredLogger
}

func (l *Grpc) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

func (l *Grpc) Infoln(args ...interface{}) {
	l.SugaredLogger.Info(args, "\n")
}

func (l *Grpc) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

func (l *Grpc) Warning(args ...interface{}) {
	l.SugaredLogger.Warn(args)
}

func (l *Grpc) Warningln(args ...interface{}) {
	l.SugaredLogger.Warn(args, "\n")
}

func (l *Grpc) Warningf(format string, args ...interface{}) {
	l.SugaredLogger.Warnf(format, args...)
}

func (l *Grpc) Error(args ...interface{}) {
	l.SugaredLogger.Error(args)
}

func (l *Grpc) Errorln(args ...interface{}) {
	l.SugaredLogger.Error(args, "\n")
}

func (l *Grpc) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf("format", args...)
}

func (l *Grpc) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args)
}

func (l *Grpc) Fatalln(args ...interface{}) {
	l.SugaredLogger.Fatal(args, "\n")
}

func (l *Grpc) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Fatalf(format, args...)
}

func (l *Grpc) V(int) bool {
	return true
}
