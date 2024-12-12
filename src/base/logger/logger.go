package logger

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	DEBUG   = 0b00001
	INFO    = 0b00010
	WARNING = 0b00100
	ERROR   = 0b01000
	FATAL   = 0b10000

	ALL  = DEBUG | INFO | WARNING | ERROR | FATAL
	NONE = 0
)

type dnxLogger struct {
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger

	LogToFile    bool
	LogToConsole bool
	LogOptions   int
}

var dnxLoggerInstance *dnxLogger

func Init() {
	dnxLoggerInstance = &dnxLogger{
		LogToFile:    true,
		LogToConsole: true,
		LogOptions:   ALL,

		DebugLogger:   log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
		InfoLogger:    log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
		WarningLogger: log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Lshortfile),
		ErrorLogger:   log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile),
		FatalLogger:   log.New(os.Stderr, "[FATAL] ", log.LstdFlags|log.Lshortfile),
	}

	EnvInit()
}

func EnvInit() {
	minLogLevel, existMinLevel := os.LookupEnv("DNX_LOG_MIN_LEVEL")
	disableLevels, existDisableLevels := os.LookupEnv("DNX_LOG_DISABLE_LEVELS")
	logConsole, existLogConsole := os.LookupEnv("DNX_LOG_CONSOLE")
	logFile, existLogFile := os.LookupEnv("DNX_LOG_FILE")

	if existMinLevel {
		SetMinLogLevel(LogLevelValue(minLogLevel))
	}
	if existDisableLevels {
		levels := strings.Split(disableLevels, "|")
		options := 0
		for _, level := range levels {
			level = strings.TrimSpace(level)
			options |= LogLevelValue(level)
		}

		DisableLogOptions(options)
	}
	if existLogConsole {
		b, err := strconv.ParseBool(logConsole)
		if err != nil {
			Warning("Failed to parse DNX_LOG_CONSOLE value")
		} else {
			SetLogToConsole(b)
		}
	}
	if existLogFile {
		b, err := strconv.ParseBool(logFile)
		if err != nil {
			Warning("Failed to parse DNX_LOG_FILE value")
		} else {
			SetLogToFile(b)
		}
	}
}

func LogsToFile() bool {
	return dnxLoggerInstance.LogToFile
}
func SetLogToFile(value bool) {
	Info("File logging set to", value)
	dnxLoggerInstance.LogToFile = value
}

func LogsToConsole() bool {
	return dnxLoggerInstance.LogToConsole
}
func SetLogToConsole(value bool) {
	Info("Console logging set to", value)
	dnxLoggerInstance.LogToConsole = value
}

func LogOptions() int {
	return dnxLoggerInstance.LogOptions
}
func LogOptionsHas(option int) bool {
	return LogOptions()&option == option
}
func SetLogOptions(options int) {
	if options < NONE || options > ALL {
		Warning("Invalid logging options")
		return
	} else if options == ALL {
		Info("Logging options set to ALL")
		return
	} else if options == NONE {
		Info("Logging options set to NONE")
		return
	}

	msg := "Logging options set to: "

	if LogOptionsHas(DEBUG) {
		msg += "| DEBUG |"
	}
	if LogOptionsHas(INFO) {
		msg += "| INFO |"
	}
	if LogOptionsHas(WARNING) {
		msg += "| WARNING |"
	}
	if LogOptionsHas(ERROR) {
		msg += "| ERROR |"
	}
	if LogOptionsHas(FATAL) {
		msg += "| FATAL |"
	}

	Info(msg)

	dnxLoggerInstance.LogOptions = options
}
func EnableLogOptions(options int) {
	if options < NONE || options > ALL {
		Warning("Invalid logging option")
		return
	}

	var msg string

	if options&DEBUG == DEBUG {
		msg += "| DEBUG |"
	}
	if options&INFO == INFO {
		msg += "| INFO |"
	}
	if options&WARNING == WARNING {
		msg += "| WARNING |"
	}
	if options&ERROR == ERROR {
		msg += "| ERROR |"
	}
	if options&FATAL == FATAL {
		msg += "| FATAL |"
	}

	Info("Enabled logging options: ", msg)
	dnxLoggerInstance.LogOptions |= options
}
func DisableLogOptions(options int) {
	if options < NONE || options > ALL {
		Warning("Invalid logging option")
		return
	}

	var msg string

	if options&DEBUG == DEBUG {
		msg += "| DEBUG |"
	}
	if options&INFO == INFO {
		msg += "| INFO |"
	}
	if options&WARNING == WARNING {
		msg += "| WARNING |"
	}
	if options&ERROR == ERROR {
		msg += "| ERROR |"
	}
	if options&FATAL == FATAL {
		msg += "| FATAL |"
	}

	Info("Disabled logging options: ", msg)
	dnxLoggerInstance.LogOptions &= ^options
}
func SetMinLogLevel(level int) {
	if level < NONE || level > ALL {
		Warning("Invalid logging level")
		return
	}

	var msg string

	switch level {
	case DEBUG:
		msg += "| DEBUG |"
		fallthrough
	case INFO:
		msg += "| INFO |"
		fallthrough
	case WARNING:
		msg += "| WARNING |"
		fallthrough
	case ERROR:
		msg += "| ERROR |"
		fallthrough
	case FATAL:
		msg += "| FATAL |"
		fallthrough
	case ALL:
		msg = "| ALL |"
	case NONE:
		msg = "| NONE |"
	}

	Info("Minimum logging level set to: ", msg)
	SetLogOptions(level)
}
func LogLevelValue(level string) int {
	value := LogOptions()

	switch level {
	case "DEBUG":
		value = DEBUG
	case "INFO":
		value = INFO
	case "WARNING":
		value = WARNING
	case "ERROR":
		value = ERROR
	case "FATAL":
		value = FATAL
	case "ALL":
		value = ALL
	case "NONE":
		value = NONE
	default:
		Warning("Invalid logging level")
	}

	return value
}

func canLogWith(logger *log.Logger) bool {
	if LogOptionsHas(ALL) {
		return true
	} else if LogOptionsHas(NONE) {
		return false
	}

	if logger == dnxLoggerInstance.DebugLogger {
		return LogOptionsHas(DEBUG)
	} else if logger == dnxLoggerInstance.InfoLogger {
		return LogOptionsHas(INFO)
	} else if logger == dnxLoggerInstance.WarningLogger {
		return LogOptionsHas(WARNING)
	} else if logger == dnxLoggerInstance.ErrorLogger {
		return LogOptionsHas(ERROR)
	} else if logger == dnxLoggerInstance.FatalLogger {
		return LogOptionsHas(FATAL)
	}

	return false
}

func writeToFile(prefix string, v ...interface{}) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logError(false, "Failed to open log file")
		return
	}
	defer file.Close()

	logger := log.New(file, prefix, log.LstdFlags|log.Lshortfile)
	logger.Println(v...)
}

func logWith(logger *log.Logger, ForceWriteFile bool, v ...interface{}) {
	if !canLogWith(logger) {
		return
	}

	if dnxLoggerInstance.LogToConsole {
		logger.Println(v...)
	}

	if ForceWriteFile || dnxLoggerInstance.LogToFile {
		writeToFile(logger.Prefix(), v...)
	}
}

func logDebug(writeFile bool, v ...interface{}) {
	logWith(dnxLoggerInstance.DebugLogger, writeFile, v...)
}
func logInfo(writeFile bool, v ...interface{}) {
	logWith(dnxLoggerInstance.InfoLogger, writeFile, v...)
}
func logWarning(writeFile bool, v ...interface{}) {
	logWith(dnxLoggerInstance.WarningLogger, writeFile, v...)
}
func logError(writeFile bool, v ...interface{}) {
	logWith(dnxLoggerInstance.ErrorLogger, writeFile, v...)
}
func logFatal(writeFile bool, v ...interface{}) {
	logWith(dnxLoggerInstance.FatalLogger, writeFile, v...)
	os.Exit(1)
}

func Debug(v ...interface{}) {
	logDebug(false, v...)
}
func Info(v ...interface{}) {
	logInfo(false, v...)
}
func Warning(v ...interface{}) {
	logWarning(false, v...)
}
func Error(v ...interface{}) {
	logError(false, v...)
}
func Fatal(v ...interface{}) {
	logFatal(false, v...)
	os.Exit(1)
}
