package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func Prefix() string {
	return "logger-" + time.Now().Format("2006-01-02")
}

func InitLog() *zap.SugaredLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
	sugarLogger := logger.Sugar()
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.TimeKey = "timestamp"
	loggerConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000Z07:00")
	loggerConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	loggerConfig.FunctionKey = "func"
	return zapcore.NewJSONEncoder(loggerConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logFile := "logger/" + Prefix() + ".log"
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename: logFile,
		Compress: false,
	})
	return ws
}

// only message
func Info(msg string)    { sugar := InitLog(); sugar.Info(msg) }
func Warning(msg string) { sugar := InitLog(); sugar.Warn(msg) }
func Error(msg string)   { sugar := InitLog(); sugar.Error(msg) }
func Fatal(msg string)   { sugar := InitLog(); sugar.Fatal(msg) }
func Panic(msg string)   { sugar := InitLog(); sugar.Panic(msg) }

// with data
func InfoW(msg string, data interface{})    { sugar := InitLog(); sugar.Infow(msg, "data", data) }
func WarningW(msg string, data interface{}) { sugar := InitLog(); sugar.Warnw(msg, "data", data) }
func ErrorW(msg string, data interface{})   { sugar := InitLog(); sugar.Errorw(msg, "data", data) }
func FatalW(msg string, data interface{})   { sugar := InitLog(); sugar.Fatalw(msg, "data", data) }
func PanicW(msg string, data interface{})   { sugar := InitLog(); sugar.Panicw(msg, "data", data) }
