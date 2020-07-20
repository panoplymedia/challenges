package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

var logLevels = map[string]logLevel{
	"DEBUG": logLevel(log.DEBUG),
	"INFO":  logLevel(log.INFO),
	"WARN":  logLevel(log.WARN),
	"ERROR": logLevel(log.ERROR),
	"OFF":   logLevel(log.OFF),
}

var logLevelStrings = map[logLevel]string{
	logLevel(log.DEBUG): "DEBUG",
	logLevel(log.INFO):  "INFO",
	logLevel(log.WARN):  "WARN",
	logLevel(log.ERROR): "ERROR",
	logLevel(log.OFF):   "OFF",
}

// logLevel is the output level of the logs: DEBUG, INFO, WARN, ERROR, or OFF.
type logLevel log.Lvl

// Impements set's the logLevel from a string: DEBUG, INFO, WARN, ERROR, or OFF.
func (l *logLevel) Set(s string) error {
	level, ok := logLevels[s]
	if !ok {
		return errors.Errorf("'%s' is not a valid log level. Use DEBUG, INFO, WARN, ERROR, or OFF", s)
	}

	*l = level
	return nil
}

func (l logLevel) Get() interface{} {
	return l
}

func (l logLevel) String() string {
	s, ok := logLevelStrings[l]
	if !ok {
		return fmt.Sprintf("Unknown log level: %d", l)
	}

	return s
}
