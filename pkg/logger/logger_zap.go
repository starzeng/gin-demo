package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var IsDev = true

var log *zap.SugaredLogger

func init() {
	var _ zapcore.Encoder

	if IsDev {
		encCfg := zap.NewDevelopmentEncoderConfig()
		encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		_ = zapcore.NewConsoleEncoder(encCfg)
	} else {
		encCfg := zap.NewProductionEncoderConfig()
		_ = zapcore.NewJSONEncoder(encCfg)
	}

	logger, _ := zap.NewProduction()
	log = logger.Sugar()

}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Infof(msg)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	log.Fatalw(msg)
}
