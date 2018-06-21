package logo

import (
	"bytes"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	debugLog = log.New(os.Stdout, "DEBUG:   ", 0)
	infoLog = log.New(os.Stdout, "INFO:    ", 0)
	warningLog = log.New(os.Stdout, "WARNING: ", 0)
	errorLog = log.New(os.Stdout, "ERROR:   ", 0)
}

const (
	DEBUG_LEVEL   = 1
	INFO_LEVEL    = 0
	WARNING_LEVEL = -1
	ERROR_LEVEL   = -2
	DISABLED      = -3
)

var (
	infoLogChannel    chan []string
	debugLogChannel   chan []string
	warningLogChannel chan []string
	errorLogChannel   chan []string

	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	logLevel   = INFO_LEVEL
)

// listenLogChannel adds all elements of the string slice into a byte buffer,
// afterwards the content of the buffer will be written
// into the specified logger.
func listenLogChannel(logWriter *log.Logger, channel chan []string) {
	var logBuffer bytes.Buffer

	for logItems := range channel {
		logBuffer.WriteString(time.Now().Format(time.Stamp))
		logBuffer.WriteString(" ")

		for _, logItem := range logItems {
			logBuffer.WriteString(logItem)
		}

		logWriter.Println(logBuffer.String())
		logBuffer.Reset()
	}
}

// InitNewLogger sets the io.Writer interface as an output for
// the loggers and spawns a goroutine for each of the
// available loggers.
func InitNewLogger(outInterface io.Writer, level int) {
	infoLogChannel = make(chan []string, 2<<4)
	debugLogChannel = make(chan []string, 2<<8)
	warningLogChannel = make(chan []string, 2<<4)
	errorLogChannel = make(chan []string, 2<<4)

	logLevel = level

	if logLevel != DISABLED {
		debugLog.SetOutput(outInterface)
		infoLog.SetOutput(outInterface)
		warningLog.SetOutput(outInterface)
		errorLog.SetOutput(outInterface)

		go listenLogChannel(debugLog, debugLogChannel)
		go listenLogChannel(errorLog, errorLogChannel)
		go listenLogChannel(warningLog, warningLogChannel)
		go listenLogChannel(infoLog, infoLogChannel)
	}
}

// DebugSync writes a message to the debug log in a
// synchronous manner, if debug level enabled.
func DebugSync(message string) {
	if DEBUG_LEVEL <= logLevel {
		debugLog.Println(message)
	}
}

// Sends the given strings to the debug log chanel.
func Debug(message ...string) {
	if DEBUG_LEVEL <= logLevel {
		debugLogChannel <- message
	}
}

// InfoSync writes a message to the info log in a
// synchronous manner, if info level enabled.
func InfoSync(message string) {
	if INFO_LEVEL <= logLevel {
		infoLog.Println(message)
	}
}

// Info sends the given string arguments to
// the info log chanel.
func Info(message ...string) {
	if INFO_LEVEL <= logLevel {
		infoLogChannel <- message
	}
}

// WarningSync writes a message to the warning log
// in a synchronous manner, if warning level enabled.
func WarningSync(message string) {
	if WARNING_LEVEL <= logLevel {
		warningLog.Println(message)
	}
}

// Warning sends the given string arguments to the
// warning log chanel.
func Warning(message ...string) {
	if WARNING_LEVEL <= logLevel {
		warningLogChannel <- message
	}
}

// ErrorSync writes a message to the error log
// in a synchronous manner, if error level enabled.
func ErrorSync(message string) {
	if ERROR_LEVEL <= logLevel {
		errorLog.Println(message)
	}
}

// Error sends the given string arguments to the
// error log chanel.
func Error(message ...string) {
	if ERROR_LEVEL <= logLevel {
		errorLogChannel <- message
	}
}
