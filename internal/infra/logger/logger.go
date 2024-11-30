package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct{
	log *zap.Logger
}

func NewLogger() *Logger{
	logConfiguration := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey: "level",
			TimeKey: "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ := logConfiguration.Build()

	return &Logger{
		log: log,
	}
}

func (l *Logger) Info(message string,tags ...zap.Field){
	l.log.Info(message,tags...)
	l.log.Sync()
}

func (l *Logger) Error(message string,err error, tags ...zap.Field){
	tags = append(tags,zap.NamedError("error",err))
	l.log.Error(message,tags...)
	l.log.Sync()
}