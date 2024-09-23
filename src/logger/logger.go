package logger

import (
	"log"
	"os"
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

var (
	// debugLogger is a logger for debug logs
	debugLogger *log.Logger
	// infoLogger is a logger for info logs
	infoLogger *log.Logger
	// warningLogger is a logger for warning logs
	warningLogger *log.Logger
	// errorLogger is a logger for error logs
	errorLogger *log.Logger
	// fatalLogger is a logger for fatal logs
	fatalLogger *log.Logger

	logToFile    bool
	logToConsole bool
	logOptions   int
)

func init() {
	logToFile = true
	logToConsole = true
	logOptions = ALL

	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	fatalLogger = log.New(os.Stderr, "[FATAL] ", log.LstdFlags|log.Lshortfile)
}

func LogsToFile() bool {
	return logToFile
}
func SetLogToFile(value bool) {
	Info("File logging set to", value)
	logToFile = value
}

func LogsToConsole() bool {
	return logToConsole
}
func SetLogToConsole(value bool) {
	Info("Console logging set to", value)
	logToConsole = value
}

func LogOptionsSet() int {
	return logOptions
}
func LogOptionsHas(option int) bool {
	return logOptions&option == option
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

	Info(msg)

	logOptions = options
}
func EnableLogOption(option int) {
	if option < NONE || option > ALL {
		Warning("Invalid logging option")
		return
	}

	var value string

	if option&DEBUG == DEBUG {
		value = "DEBUG"
	} else if option&INFO == INFO {
		value = "INFO"
	} else if option&WARNING == WARNING {
		value = "WARNING"
	} else if option&ERROR == ERROR {
		value = "ERROR"
	} else if option&FATAL == FATAL {
		value = "FATAL"
	}

	Info("Enabled logging option: ", value)
	logOptions |= option
}
func DisableLogOption(option int) {
	if option < NONE || option > ALL {
		Warning("Invalid logging option")
		return
	}

	var value string

	if option&DEBUG == DEBUG {
		value = "DEBUG"
	} else if option&INFO == INFO {
		value = "INFO"
	} else if option&WARNING == WARNING {
		value = "WARNING"
	} else if option&ERROR == ERROR {
		value = "ERROR"
	} else if option&FATAL == FATAL {
		value = "FATAL"
	}

	Info("Disabled logging option: ", value)
	logOptions &= ^option
}

func writeToFile(prefix string, v ...interface{}) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		error(false, "Failed to open log file")
		return
	}
	defer file.Close()

	logger := log.New(file, prefix, log.LstdFlags|log.Lshortfile)
	logger.Println(v...)
}
func canLogWith(logger *log.Logger) bool {
	if logger == debugLogger && !LogOptionsHas(DEBUG) {
		return false
	} else if logger == infoLogger && !LogOptionsHas(INFO) {
		return false
	} else if logger == warningLogger && !LogOptionsHas(WARNING) {
		return false
	} else if logger == errorLogger && !LogOptionsHas(ERROR) {
		return false
	} else if logger == fatalLogger && !LogOptionsHas(FATAL) {
		return false
	}

	return true
}

func logWith(logger *log.Logger, ForceWriteFile bool, v ...interface{}) {
	if !canLogWith(logger) {
		return
	}

	if logToConsole {
		logger.Println(v...)
	}

	if ForceWriteFile || logToFile {
		writeToFile(logger.Prefix(), v...)
	}
}

func Debug(v ...interface{}) {
	logWith(debugLogger, true, v...)
}
func Info(v ...interface{}) {
	logWith(infoLogger, true, v...)
}
func Warning(v ...interface{}) {
	logWith(warningLogger, true, v...)
}
func Error(v ...interface{}) {
	logWith(errorLogger, true, v...)
}
func Fatal(v ...interface{}) {
	logWith(fatalLogger, true, v...)
	os.Exit(1)
}

func debug(writeFile bool, v ...interface{}) {
	logWith(debugLogger, writeFile, v...)
}
func info(writeFile bool, v ...interface{}) {
	logWith(infoLogger, writeFile, v...)
}
func warning(writeFile bool, v ...interface{}) {
	logWith(warningLogger, writeFile, v...)
}
func error(writeFile bool, v ...interface{}) {
	logWith(errorLogger, writeFile, v...)
}
func fatal(writeFile bool, v ...interface{}) {
	logWith(fatalLogger, writeFile, v...)
	os.Exit(1)
}
