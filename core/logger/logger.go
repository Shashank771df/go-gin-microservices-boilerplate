package logger

import (
	"strings"
)

const (
	DEFAULT_LOGGER = "DEAFULT_LOGGER"
)

const (
	RESET   = "\x1b[0m"
	MAGENTA = "\x1b[95m"
	YELLOW  = "\x1b[33m"
	GREEN   = "\x1b[32m"
	CYAN    = "\x1b[96m"
	RED     = "\x1b[31m"
)

const (
	ERROR   = "ERROR"
	WARNING = "WARNING"
	INFO    = "INFO"
	DEBUG   = "DEBUG"
)

const (
	ERROR_VALUE   = 1
	WARNING_VALUE = 2
	INFO_VALUE    = 3
	DEBUG_VALUE   = 4
)

var logLevels map[string]byte = map[string]byte{
	ERROR:   ERROR_VALUE,
	WARNING: WARNING_VALUE,
	INFO:    INFO_VALUE,
	DEBUG:   DEBUG_VALUE,
}

type LogInfo struct {
	Key   string
	Value string
	Extra map[string]interface{}
}

type LogEnabled struct {
	Warning bool
	Debug   bool
	Info    bool
	Error   bool
}

type LogProps struct {
	ProjectID      string
	Name           string
	MaxFiles       uint16
	FilePrefix     string
	FileSize       string
	Path           string
	ConsoleEnabled bool
	FileEnabled    bool
	MaxQueue       uint16
	LogLevel       string
}

type Logger interface {
	Error(log LogInfo)
	Warn(log LogInfo)
	Info(log LogInfo)
	Fatal(log LogInfo)
	Debug(log LogInfo)
	Init(log LogProps)
	SetLogLevel(value string)
}

func New(logger string, props LogProps) Logger {
	var item Logger

	switch logger {
	case DEFAULT_LOGGER:
		item = &DefaultLogger{}
	default:
		item = &DefaultLogger{}
	}

	item.Init(props)

	return item
}

func (o *LogEnabled) SetLogLevel(data string) {
	key := strings.ToUpper(data)

	var level byte
	var ok bool

	if level, ok = logLevels[key]; !ok {
		level = logLevels[ERROR]
	}

	o.Error = level >= ERROR_VALUE
	o.Warning = level >= WARNING_VALUE
	o.Info = level >= INFO_VALUE
	o.Debug = level >= DEBUG_VALUE
}
