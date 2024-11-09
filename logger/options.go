package logger

import (
	"io"
)

type Options struct {
	Level     *LogLevel
	Prefix    *string
	Colorized *bool
	Writer    io.Writer
}

type LogLevel uint8

var LogLevels = struct {
	Fatal LogLevel
	Error LogLevel
	Warn  LogLevel
	Info  LogLevel
	Debug LogLevel
}{
	Fatal: 0,
	Error: 1,
	Warn:  2,
	Info:  3,
	Debug: 4,
}
