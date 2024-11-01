package logger

type loggerOptions struct {
	level       LogLevel
	prettyPrint bool
	prefix      string
}

type Options struct {
	Level  *LogLevel
	Pretty *bool
	Prefix *string
}
type LogLevel uint8

const (
	fatalFlag LogLevel = 0x1 << 0
	errorFlag LogLevel = 0x1 << 1
	warnFlag  LogLevel = 0x1 << 2
	infoFlag  LogLevel = 0x1 << 3
	debugFlag LogLevel = 0x1 << 4

	fatalLevel LogLevel = fatalFlag
	errorLevel LogLevel = fatalLevel | errorFlag
	warnLevel  LogLevel = errorLevel | warnFlag
	infoLevel  LogLevel = warnLevel | infoFlag
	debugLevel LogLevel = infoLevel | debugFlag
)

var LogLevels = struct {
	Debug LogLevel
	Info  LogLevel
	Warn  LogLevel
	Error LogLevel
	Fatal LogLevel
}{
	Debug: debugLevel,
	Info:  infoLevel,
	Warn:  warnLevel,
	Error: errorLevel,
	Fatal: fatalLevel,
}
