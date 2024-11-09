package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	level        LogLevel
	lvlSubStrMap map[LogLevel]string
	logger       *log.Logger
}

const (
	// Colors
	resetColor = "\033[0m"
	debugColor = "\033[34m"
	infoColor  = "\033[32m"
	warnColor  = "\033[33m"
	errorColor = "\033[31m"
	fatalColor = "\033[41m"
)

var instance *Logger

func init() {
	instance = New()
}

func lvlSubStringMap(colorized bool) map[LogLevel]string {
	result := make(map[LogLevel]string)
	result[LogLevels.Debug] = paintLvlStr(colorized, "DEBUG", debugColor)
	result[LogLevels.Info] = paintLvlStr(colorized, "INFO", infoColor)
	result[LogLevels.Warn] = paintLvlStr(colorized, "WARNING", warnColor)
	result[LogLevels.Error] = paintLvlStr(colorized, "ERROR", errorColor)
	result[LogLevels.Fatal] = paintLvlStr(colorized, "FATAL", fatalColor)
	return result
}

func paintLvlStr(colorized bool, level string, color string) string {
	if colorized {
		return fmt.Sprintf("%s%s%s", color, level, resetColor)
	}
	return level
}

func (l *Logger) printLog(logLevel LogLevel, msg ...interface{}) {
	log.Printf("[%s] %v", l.lvlSubStrMap[logLevel], fmt.Sprint(msg...))
}

func (l *Logger) Debug(msg ...interface{}) {
	if l.level >= LogLevels.Debug {
		l.printLog(LogLevels.Debug, msg...)
	}
}

func (l *Logger) DebugF(format string, msg ...interface{}) {
	if l.level >= LogLevels.Debug {
		l.printLog(LogLevels.Debug, fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Info(msg ...interface{}) {
	if l.level >= LogLevels.Info {
		l.printLog(LogLevels.Info, msg...)
	}
}

func (l *Logger) InfoF(format string, msg ...interface{}) {
	if l.level >= LogLevels.Info {
		l.printLog(LogLevels.Info, fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Warn(msg ...interface{}) {
	if l.level >= LogLevels.Warn {
		l.printLog(LogLevels.Warn, msg...)
	}
}

func (l *Logger) WarnF(format string, msg ...interface{}) {
	if l.level >= LogLevels.Warn {
		l.printLog(LogLevels.Warn, fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Error(msg ...interface{}) {
	if l.level >= LogLevels.Error {
		l.printLog(LogLevels.Error, msg...)
	}
}

func (l *Logger) ErrorF(format string, msg ...interface{}) {
	if l.level >= LogLevels.Error {
		l.printLog(LogLevels.Error, fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Fatal(msg ...interface{}) {
	l.printLog(LogLevels.Fatal, msg...)
}

func (l *Logger) FatalF(format string, msg ...interface{}) {
	l.printLog(LogLevels.Fatal, fmt.Sprintf(format, msg...))
}

func New() *Logger {
	l := log.New(os.Stdout, "", log.LUTC|log.Ldate|log.Lmicroseconds)
	return &Logger{
		level:        LogLevels.Info,
		lvlSubStrMap: lvlSubStringMap(false),
		logger:       l,
	}
}

func Configure(opts Options) {
	if opts.Level != nil {
		instance.level = *opts.Level
	}
	if opts.Colorized != nil {
		instance.lvlSubStrMap = lvlSubStringMap(*opts.Colorized)
	}
	if opts.Prefix != nil {
		instance.logger.SetPrefix(*opts.Prefix)
	}
	if opts.Writer != nil {
		instance.logger.SetOutput(opts.Writer)
	}
}

func Debug(msg ...interface{}) {
	instance.Debug(msg...)
}

func DebugF(format string, msg ...interface{}) {
	instance.DebugF(format, msg...)
}

func Info(msg ...interface{}) {
	instance.Info(msg...)
}

func InfoF(format string, msg ...interface{}) {
	instance.InfoF(format, msg...)
}

func Warn(msg ...interface{}) {
	instance.Warn(msg...)
}

func WarnF(format string, msg ...interface{}) {
	instance.WarnF(format, msg...)
}

func Error(msg ...interface{}) {
	instance.Error(msg...)
}

func ErrorF(format string, msg ...interface{}) {
	instance.ErrorF(format, msg...)
}

func Fatal(msg ...interface{}) {
	instance.Fatal(msg...)
}

func FatalF(format string, msg ...interface{}) {
	instance.FatalF(format, msg...)
}
