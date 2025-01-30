package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func GetLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(levelDebug string) *LoggerZap {
	var level zapcore.Level
	switch levelDebug {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}

	hook := lumberjack.Logger{
		Filename:   "./logs/app.logger",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	var encoder zapcore.Encoder
	switch level {
	case zapcore.DebugLevel:
		encoder = zapcore.NewConsoleEncoder(GetEncoderLog())
	case zapcore.InfoLevel:
		encoder = zapcore.NewJSONEncoder(GetEncoderLog())
	}

	core := zapcore.NewCore(
		encoder,
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func GetEncoderLog() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encoderConfig.TimeKey = "time"

	encoderConfig.CallerKey = "caller"

	encoderConfig.LevelKey = "level"

	encoderConfig.MessageKey = "message"

	encoderConfig.EncodeLevel = CustomLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoderConfig.EncodeTime = SyslogTimeEncoder

	return encoderConfig
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
