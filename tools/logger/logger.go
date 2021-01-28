package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         zap.NewProductionConfig().Encoding,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "lv",
			TimeKey:      "ts",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = cfg.Build(zap.AddCallerSkip(1)); err != nil {
		panic(err)
	}
}

// Debug log
func Debug(msg string, fields ...zap.Field) {
	defer log.Sync()

	log.Debug(msg, fields...)
}

// Info log
func Info(msg string, fields ...zap.Field) {
	defer log.Sync()

	log.Info(msg, fields...)
}

// Warn log
func Warn(msg string, fields ...zap.Field) {
	defer log.Sync()

	log.Warn(msg, fields...)
}

// Fatal log
func Fatal(msg string, fields ...zap.Field) {
	defer log.Sync()

	log.Fatal(msg, fields...)
}

// Error log
func Error(msg string, err error, fields ...zap.Field) {
	defer log.Sync()
	msg = fmt.Sprintf("%s - %v", msg, err)
	log.Error(msg, fields...)
}
