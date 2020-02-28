package logger

import (
	"context"
	"fmt"

	"github.com/DeNA/aelog"
)

func Debugf(ctx context.Context, format string, a ...interface{}) {
	aelog.Debugf(ctx, format, a...)
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	aelog.Infof(ctx, format, a...)
}

func Warningf(ctx context.Context, format string, a ...interface{}) {
	aelog.Warningf(ctx, format, a...)
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	aelog.Errorf(ctx, format, a...)
}

func Criticalf(ctx context.Context, format string, a ...interface{}) {
	aelog.Criticalf(ctx, format, a...)
}

func Debug(ctx context.Context, a ...interface{}) {
	aelog.Debugf(ctx, fmt.Sprint(a...))
}

func Info(ctx context.Context, a ...interface{}) {
	aelog.Infof(ctx, fmt.Sprint(a...))
}

func Warning(ctx context.Context, a ...interface{}) {
	aelog.Warningf(ctx, fmt.Sprint(a...))
}

func Error(ctx context.Context, a ...interface{}) {
	aelog.Errorf(ctx, fmt.Sprint(a...))
}

func Critical(ctx context.Context, a ...interface{}) {
	aelog.Criticalf(ctx, fmt.Sprint(a...))
}
