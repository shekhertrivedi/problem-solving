package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type level uint8

const (
	fatallevel level = iota
	errorlevel
	warninglevel
	infolevel
	verboselevel
)

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	Verbose(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

type logger struct {
	name string
}

func (log *logger) Verbose(v ...interface{}) {
	log.output(verboselevel, v...)
}

func (log *logger) Info(v ...interface{}) {
	log.output(infolevel, v...)
}

func (log *logger) Warning(v ...interface{}) {
	log.output(warninglevel, v...)
}

func (log *logger) Error(v ...interface{}) {
	log.output(errorlevel, v...)
}

func (log *logger) Fatal(v ...interface{}) {
	log.output(fatallevel, v...)
}

func (log *logger) output(l level, v ...interface{}) {
	if level, ok := componentsLevels[log.name]; ok {
		if level < l {
			return
		}
	} else {
		if defaultLevel < l {
			return
		}
	}

	mu.Lock()
	fmt.Println(l, log.name, "\t-", fmt.Sprint(v...))
	mu.Unlock()
}

func (l level) String() (name string) {
	switch l {
	case verboselevel:
		name = "[VERBOSE]"
	case infolevel:
		name = "[INFO]   "
	case warninglevel:
		name = "[WARNING]"
	case errorlevel:
		name = "[ERROR]  "
	case fatallevel:
		name = "[FATAL]  "
	default:
		name = "<UNKNOWN>"
	}
	return
}

// NewLogger returns named log object with implemented Logger interface
func NewLogger(name string) (log Logger) {
	log = &logger{name: name}
	return
}

// PrintSettings prints current log levels setting per named log to stdout
func PrintSettings() {
	fmt.Println("Default :", defaultLevel)
	for k, v := range componentsLevels {
		fmt.Println(k, ":", v)
	}
}

var (
	mu               sync.Mutex
	defaultLevel     = fatallevel
	componentsLevels map[string]level
)

func strtolevel(s string) (level, error) {
	switch s {
	case "FATAL":
		return fatallevel, nil
	case "ERROR":
		return errorlevel, nil
	case "WARNING":
		return warninglevel, nil
	case "INFO":
		return infolevel, nil
	case "VERBOSE":
		return verboselevel, nil
	}
	return fatallevel, errors.New("Unknown logs level " + s)
}

func init() {
	componentsLevels = make(map[string]level)
	if levels := os.Getenv("LOGLEVELS"); levels != "" {
		levels := strings.Split(levels, ";")
		for _, strlevel := range levels {
			strlevel := strings.Split(strlevel, ":")
			switch len(strlevel) {
			case 1:
				if level, err := strtolevel(strings.ToUpper(strlevel[0])); err == nil {
					defaultLevel = level
				}
			case 2:
				if level, err := strtolevel(strings.ToUpper(strlevel[1])); err == nil {
					componentsLevels[strlevel[0]] = level
				}
			}
		}
	}
}
