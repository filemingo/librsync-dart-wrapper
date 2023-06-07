package logging

import (
	"sync"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const tag = "librsync:go"

type Logger interface {
	D(TAG, message string)
	E(TAG, message string)
	V(TAG, message string)
	I(TAG, message string)
	W(TAG, message string)
}

type dummyLogger int

func (d dummyLogger) D(TAG, message string) {
}
func (d dummyLogger) E(TAG, message string) {
}
func (d dummyLogger) V(TAG, message string) {
}
func (d dummyLogger) I(TAG, message string) {
}

func (d dummyLogger) W(TAG, message string) {
}

var DummyLoggerInstance dummyLogger

var loggers = map[string]Logger{}
var logMutex sync.Mutex

func AddLogger(logger Logger, id ...string) func() {
	var finalID string
	if len(id) > 0 {
		finalID = id[0]
	} else {
		finalID = uuid.New().String()
	}
	logMutex.Lock()
	defer logMutex.Unlock()
	loggers[finalID] = logger
	return func() {
		logMutex.Lock()
		defer logMutex.Unlock()
		delete(loggers, finalID)
	}
}

type loggerHook struct{}

func (l loggerHook) Fire(entry *log.Entry) error {
	log.Trace("Log hook fired\n")
	msg := entry.Message
	logMutex.Lock()
	defer logMutex.Unlock()

	for _, Logger := range loggers {
		switch entry.Level {
		case log.DebugLevel:
			Logger.D(tag, msg)
		case log.InfoLevel:
			Logger.I(tag, msg)
		case log.WarnLevel:
			Logger.W(tag, msg)
		case log.ErrorLevel:
			fallthrough
		case log.FatalLevel:
			fallthrough
		case log.PanicLevel:
			Logger.E(tag, msg)
		}
	}
	return nil
}

func (l loggerHook) Levels() []log.Level {
	return []log.Level{
		log.DebugLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.InfoLevel,
		log.PanicLevel,
		log.WarnLevel,
	}
}

var logInitOnce sync.Once

func LogInit() {
	logInitOnce.Do(func() {
		hook := loggerHook{}
		log.AddHook(hook)
		log.Infof("Finished running LogInit()\n")
	})
}
