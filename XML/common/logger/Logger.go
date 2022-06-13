package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Logger struct {
	Logger *logrus.Logger
	file   *os.File
}

func InitializeLogger(service string, ctx context.Context) *Logger {
	logger := &Logger{}
	if err := CreateLogDir(filepath.FromSlash("../logfiles/" + service)); err != nil {
		logrus.Fatalf("Failed to create directory for log files | %v\n", err)
	}
	file := filepath.FromSlash("../logfiles/" + service + "/" + service + ".log")

	rotatingLogs, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   file,
		MaxSize:    100, //megabytes
		MaxBackups: 50,  //MaxBackups is the maximum number of old log files to retain.
		MaxAge:     14,  //days
		Level:      logrus.InfoLevel,
		Formatter: UTCFormatter{&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05Z",
			DataKey:         "data",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return frame.Function, fmt.Sprintf("%s:%d", FormatFilePath(frame.File), frame.Line)
			}}},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize rotating hook | %v\n", err)
	}
	logger.Logger = logrus.New()
	logger.Logger.AddHook(rotatingLogs)
	logger.Logger.SetReportCaller(true)
	logger.Logger.SetOutput(os.Stdout)
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, fmt.Sprintf("%s:%d", FormatFilePath(frame.File), frame.Line)
		}})
	logger.Logger.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: logger.Logger, file: logger.file}

}

func FormatFilePath(file string) any {
	arr := strings.Split(filepath.ToSlash(file), "/")
	return arr[len(arr)-1]
}

func CreateLogDir(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.MkdirAll(filePath, os.ModeDir|0777)
	}
	return nil
}

func (l *Logger) CloseLogger() {
	l.file.Close()
}

func (l *Logger) ErrorMessage(message string) {
	l.Logger.Error(message)
}

func (l *Logger) InfoMessage(message string) {
	l.Logger.Info(message)
}

func (l *Logger) FatalMessage(message string) {
	l.Logger.Fatal(message)
}

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}
