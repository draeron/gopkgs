package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logrus struct {
	*logrus.Entry
}

type formatter struct {
	*logrus.TextFormatter
}

func NewLogrus(name string) *Logrus {
	log := logrus.Logger{
		Out: os.Stdout,
		Formatter: &formatter{
			TextFormatter: &logrus.TextFormatter{
				ForceColors: true,
				CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
					return "", fmt.Sprintf("%s:%d", formatFilePath(frame.File), frame.Line)
				},
			},
		},
		Level: logrus.InfoLevel,
	}
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return &Logrus{Entry: log.WithField("logger", name)}
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func (l *Logrus) ErrorIf(err error, msgfmt string, args ...interface{}) {
	if err != nil {
		l.Errorf(msgfmt+fmt.Sprintf(": %+v", err), args...)
	}
}
