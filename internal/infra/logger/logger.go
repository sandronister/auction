package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = &zap.Logger{}

func Init() {
	logConfiguration := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log, err = logConfiguration.Build()
	if err != nil {
		panic(err)
	}
	

}

func  Info(message string,tags ...zap.Field){
	log.Info(message,tags...)
	log.Sync()
}

func  Error(message string,err error, tags ...zap.Field){
	tags = append(tags,zap.NamedError("error",err))
	log.Error(message,tags...)
	log.Sync()
}