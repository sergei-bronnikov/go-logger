package logger

import (
	"fmt"
	"log"
)

const (
	// Colors
	resetColor = "\033[0m"
	debugColor = "\033[34m"
	infoColor  = "\033[32m"
	warnColor  = "\033[33m"
	errorColor = "\033[31m"
	fatalColor = "\033[41m"
)

var options = loggerOptions{
	level:       LogLevels.Info,
	prettyPrint: false,
	prefix:      "",
}

type levelSubstring struct {
	level string
	color string
}

func (ls *levelSubstring) String() string {
	if options.prettyPrint {
		return fmt.Sprintf("%v%v%v", ls.color, ls.level, resetColor)
	}
	return ls.level
}

var levelSubstrings = struct {
	DEBUG levelSubstring
	INFO  levelSubstring
	WARN  levelSubstring
	ERROR levelSubstring
	FATAL levelSubstring
}{
	DEBUG: levelSubstring{"DEBUG", debugColor},
	INFO:  levelSubstring{"INFO", infoColor},
	WARN:  levelSubstring{"WARN", warnColor},
	ERROR: levelSubstring{"ERROR", errorColor},
	FATAL: levelSubstring{"FATAL", fatalColor},
}

func init() {
	//configureOutput()
	log.SetFlags(log.LUTC | log.Ldate | log.Lmicroseconds)
}

func Configure(opts Options) {
	if opts.Level != nil {
		options.level = *opts.Level
	}
	if opts.Pretty != nil {
		options.prettyPrint = *opts.Pretty
	}
	if opts.Prefix != nil {
		options.prefix = *opts.Prefix
	}
}

func printLog(levelSubstring string, msg ...interface{}) {
	log.Printf("[sphere] [%v] %v", levelSubstring, fmt.Sprint(msg...))
}

func Debug(msg ...interface{}) {
	if options.level&debugFlag > 0 {
		printLog(levelSubstrings.DEBUG.String(), msg...)
	}
}

func DebugF(msg string, args ...interface{}) {
	if options.level&debugFlag > 0 {
		printLog(levelSubstrings.DEBUG.String(), fmt.Sprintf(msg, args...))
	}
}

func Info(msg ...interface{}) {
	if options.level&infoFlag > 0 {
		printLog(levelSubstrings.INFO.String(), msg...)
	}
}

func InfoF(msg string, args ...interface{}) {
	if options.level&infoFlag > 0 {
		printLog(levelSubstrings.INFO.String(), fmt.Sprintf(msg, args...))
	}
}

func Warn(msg ...interface{}) {
	if options.level&warnFlag > 0 {
		printLog(levelSubstrings.WARN.String(), msg...)
	}
}

func WarnF(msg string, args ...interface{}) {
	if options.level&warnFlag > 0 {
		printLog(levelSubstrings.WARN.String(), fmt.Sprintf(msg, args...))
	}
}

func Error(msg ...interface{}) {
	if options.level&errorFlag > 0 {
		printLog(levelSubstrings.ERROR.String(), msg...)
	}
}

func ErrorF(msg string, args ...interface{}) {
	if options.level&errorFlag > 0 {
		printLog(levelSubstrings.ERROR.String(), fmt.Sprintf(msg, args...))
	}
}

func Fatal(msg ...interface{}) {
	if options.level&fatalFlag > 0 {
		printLog(levelSubstrings.FATAL.String(), msg...)
	}
}

func FatalF(msg string, args ...interface{}) {
	if options.level&fatalFlag > 0 {
		printLog(levelSubstrings.FATAL.String(), fmt.Sprintf(msg, args...))
	}
}
