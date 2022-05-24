package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Log *zap.SugaredLogger
}

func NewLogger() Logger {
	return Logger{
		Log: InitLog(),
	}
}

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
func (Logs Logger) Info(msg string)    { Logs.Log.Info(msg) }
func (Logs Logger) Warning(msg string) { Logs.Log.Warn(msg) }
func (Logs Logger) Error(msg string)   { Logs.Log.Error(msg) }
func (Logs Logger) Fatal(msg string)   { Logs.Log.Fatal(msg) }
func (Logs Logger) Panic(msg string)   { Logs.Log.Panic(msg) }

// with data
func (Logs Logger) InfoW(msg string, data interface{})    { Logs.Log.Infow(msg, "data", data) }
func (Logs Logger) WarningW(msg string, data interface{}) { Logs.Log.Warnw(msg, "data", data) }
func (Logs Logger) ErrorW(msg string, data interface{})   { Logs.Log.Errorw(msg, "data", data) }
func (Logs Logger) FatalW(msg string, data interface{})   { Logs.Log.Fatalw(msg, "data", data) }
func (Logs Logger) PanicW(msg string, data interface{})   { Logs.Log.Panicw(msg, "data", data) }
