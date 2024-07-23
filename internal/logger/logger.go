package logger

import (
	glog "github.com/jianlu8023/go-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	log = glog.NewSugaredLogger(
		&glog.Config{LogLevel: "DEBUG",
			DevelopMode: false,
		},
		glog.WithConsoleFormat(),
		glog.WithConsoleConfig(
			zapcore.EncoderConfig{
				MessageKey:     "msg",
				LevelKey:       "level",
				TimeKey:        "time",
				NameKey:        "logger",
				CallerKey:      "",
				FunctionKey:    "",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    glog.CustomCapitalStringLevelEncoder,
				EncodeTime:     glog.CustomTimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeName:     zapcore.FullNameEncoder,
			},
		),
	)
}

func GetLogger() *zap.SugaredLogger { return log }
