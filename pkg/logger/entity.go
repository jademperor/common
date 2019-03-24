package logger

import (
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// Entity ....
type Entity = logrus.Logger

// NewJSONLogger init logger to set filepath and format
func NewJSONLogger(logPath, filename, lv string) (*Entity, error) {
	return newLogger(logPath, filename, lv, jsonType)
}

// NewTextLogger ...
func NewTextLogger(logPath, filename, lv string) (*Entity, error) {
	return newLogger(logPath, filename, lv, textType)
}

type formaterType uint8

const (
	jsonType formaterType = iota + 1
	textType
)

func newLogger(logPath, filename, lv string, t formaterType) (*Entity, error) {
	logger := logrus.New()
	logger.AddHook(NewHook())

	switch t {
	case jsonType:
		logger.Formatter = &logrus.JSONFormatter{}
	case textType:
		logger.Formatter = &logrus.TextFormatter{}
	default:
		logger.Formatter = &logrus.TextFormatter{}
	}

	var (
		level logrus.Level
		err   error
	)
	if level, err = logrus.ParseLevel(lv); err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	if filename != "" {
		fd, err := ensureLogfile(logPath, filename)
		if err != nil {
			return nil, err
		}
		logger.Out = io.MultiWriter(os.Stdout, fd)
	} else {
		logger.Out = os.Stdout
	}
	return logger, nil
}

func ensureLogfile(dir, filename string) (fd *os.File, err error) {
	fd, err = os.OpenFile(path.Join(dir, filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil && os.IsNotExist(err) {
		// folder not existed
		if os.MkdirAll(dir, 0777) != nil {
			return nil, err
		}
		// reopen
		fd, err = os.OpenFile(path.Join(dir, filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		return
	}
	return
}
