package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init() {
	var err error

	config := zap.NewProductionConfig()

	config.Encoding = "json"

	Log, err = config.Build()

	if err != nil {
		panic(err)
	}
}

func Sync() {
	_ = Log.Sync()
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}
