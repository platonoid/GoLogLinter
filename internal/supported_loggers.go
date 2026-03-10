package internal

import (
	"errors"
	"fmt"
)

type LoggerInfo struct {
	PackagePath string
	Methods     map[string]struct{}
}

type SupportedLoggers struct {
	Loggers map[string]*LoggerInfo
}

func NewSupportedLoggers() *SupportedLoggers {
	return &SupportedLoggers{
		Loggers: make(map[string]*LoggerInfo),
	}
}

func (l *SupportedLoggers) AddLogger(loggerName string, pkgPath string) error {
	if loggerName == "" {
		return errors.New("logger loggerName cannot be empty")
	}

	if pkgPath == "" {
		return errors.New("package path cannot be empty")
	}

	if _, exists := l.Loggers[loggerName]; exists {
		return fmt.Errorf("logger '%s' already exists", loggerName)
	}

	l.Loggers[loggerName] = &LoggerInfo{
		PackagePath: pkgPath,
		Methods:     make(map[string]struct{}),
	}
	return nil
}

func (l *SupportedLoggers) AddLoggerMethod(loggerName string, method string) error {
	if loggerName == "" {
		return errors.New("logger Name cannot be empty")
	}

	if _, exists := l.Loggers[loggerName]; !exists {
		return fmt.Errorf("logger '%s' does not exist", loggerName)
	}

	if method == "" {
		return errors.New("method Name cannot be empty")
	}

	l.Loggers[loggerName].Methods[method] = struct{}{}
	return nil
}

func (l *SupportedLoggers) IsLoggerCall(loggerName string, pkgPath string, method string) bool {
	logger, exists := l.Loggers[loggerName]
	if !exists {
		return false
	}

	if logger.PackagePath != pkgPath {
		return false
	}

	_, ok := logger.Methods[method]
	return ok
}
