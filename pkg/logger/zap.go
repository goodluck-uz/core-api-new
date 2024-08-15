package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Int ..
	Int = zap.Int
	// String ...
	String = zap.String
	// Error ...
	Error = zap.Error
	// Bool ...
	Bool = zap.Bool
	// Any ...
	Any = zap.Any
)

type Field = zapcore.Field

func newZapLogger(namespace, level, logFile string) *zap.Logger {
	globalLevel := parseLevel(level)

	// Setup the log file writer
	logFileWriter := getLogFileWriter(logFile)
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, logFileWriter, globalLevel),
		zapcore.NewCore(getConsoleEncoder(), zapcore.Lock(os.Stdout), globalLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Named(namespace)
	zap.RedirectStdLog(logger)

	return logger
}
func parseLevel(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelDPanic:
		return zapcore.DPanicLevel
	case LevelPanic:
		return zapcore.PanicLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
