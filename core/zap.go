package core

import (
	"fmt"
	"github.com/huwhy/commons/config"
	"github.com/huwhy/commons/util/files"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var level zapcore.Level

func NewZap(dev bool, conf *config.Zap) *zap.SugaredLogger {
	if ok := files.PathExists(conf.Dir); !ok {
		_ = os.Mkdir(conf.Dir, os.ModePerm)
	}

	appHook := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.Dir + "/application.log", // ⽇志⽂件路径
		MaxSize:    32,                            // M
		MaxBackups: 30,                            // 最多保留3个备份
		MaxAge:     7,                             //days
		Compress:   true,                          // 是否压缩 disabled by default
	})

	errHook := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.Dir + "/error.log", // ⽇志⽂件路径
		MaxSize:    32,                      // M
		MaxBackups: 30,                      // 最多保留3个备份
		MaxAge:     7,                       //days
		Compress:   true,                    // 是否压缩 disabled by default
	})

	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	alevel := zap.NewAtomicLevelAt(level)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var cores []zapcore.Core

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	infoCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(appHook),
		alevel,
	)
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	errCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(errHook),
		warnLevel,
	)
	if dev {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), alevel))
	} else {
		cores = append(cores, infoCore, errCore)
	}
	// 最后创建具体的Logger
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	logger.Info("Zap Logger init success")
	return logger.Sugar()
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(conf *config.Zap) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "traceKey",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case conf.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case conf.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case conf.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case conf.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder(conf *config.Zap) zapcore.Encoder {
	if conf.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(conf))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(conf))
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(conf *config.Zap) (core zapcore.Core) {
	writer, err := files.GetWriteSyncer(conf) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(conf), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
