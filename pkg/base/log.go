// logger using uber zap and lumberjack
package atomgo

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogLevelDebug zapcore.Level = zapcore.DebugLevel
	LogLevelInfo  zapcore.Level = zapcore.InfoLevel
	LogLevelWarn  zapcore.Level = zapcore.WarnLevel
	LogLevelError zapcore.Level = zapcore.ErrorLevel
)

func NewLogger(logPath string, logName string, logLevel zapcore.Level, devMod bool, maxSize int, maxBackups int, maxDays int, compress bool) *zap.Logger {
	logFile := filepath.Join(logPath, logName)

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize, //MB
		MaxBackups: maxBackups,
		MaxAge:     maxDays, //Days
		Compress:   compress,
	})
	core := zapcore.NewCore(
		// use NewConsoleEncoder for human readable output
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), ws),
		zap.NewAtomicLevelAt(logLevel),
	)
	var logger *zap.Logger
	if devMod {
		// development mode
		// annotate message with file name, line number, function name
		logger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		logger = zap.New(core)
	}
	zap.ReplaceGlobals(logger)
	logger.Named(logName)
	defer logger.Sync()
	return logger
}

func NewSugar(logPath string, logName string, logLevel zapcore.Level, devMod bool, maxSize int, maxBackups int, maxDays int, compress bool) *zap.SugaredLogger {
	var logger *zap.Logger
	logger = NewLogger(logPath, logName, logLevel, devMod, maxSize, maxBackups, maxDays, compress)
	return logger.Sugar()
}
