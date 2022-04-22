package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logrusInit sync.Once
)

type Logrus struct {
	*logrus.Entry
}

func NewLogrus(name string) *Logrus {
	logrusInit.Do(func() {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			// DisableLevelTruncation: false,
			PadLevelText: true,
		})
		logrus.SetOutput(os.Stdout)
	})

	return &Logrus{Entry: logrus.WithField("logger", name)}
}

func (l *Logrus) ErrorIf(err error, fmt string, args ...interface{}) {
	if err != nil {
		l.Errorf(fmt, args...)
	}
}
