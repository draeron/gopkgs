package logger

import (
	"context"
	"github.com/TheCodeTeam/goodbye"
	"go.uber.org/zap"
	"os"
)

type Logger struct {
	*zap.Logger
}

type SugaredLogger struct {
	*zap.SugaredLogger
}

func New(name string) *SugaredLogger {

	//for !flag.Parsed() {
	//	<- time.After(time.Millisecond * 50)
	//}

	l, err := zap.NewDevelopment()
	l = l.WithOptions(zap.AddStacktrace(zap.FatalLevel))

	if err != nil {
		zap.S().Panic(err)
	}

	goodbye.Register(func(ctx context.Context, s os.Signal) {
		l.Sync()
	})

	return &SugaredLogger{l.Named(name).Sugar()}
}
