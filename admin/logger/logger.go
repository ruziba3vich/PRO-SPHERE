package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Initialize initializes the global zap logger
func Initialize() {
	// Set the encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "message",
		EncodeLevel:    zapcore.CapitalLevelEncoder, // Output level as CAPITALS
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // Human-readable time
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // Short file path and line number
	}

	// Set the core configuration (Encoder + Output + Level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // Log format as JSON
		zapcore.AddSync(os.Stdout),            // Output to stdout
		zapcore.DebugLevel,                    // Log level set to debug
	)

	// Enable calling info (line number and file name) and build logger
	Logger = zap.New(core, zap.AddCaller())
}

// Sync flushes any buffered log entries
func Sync() {
	Logger.Sync()
}

func GetLogger() *zap.Logger {
	return Logger
}
