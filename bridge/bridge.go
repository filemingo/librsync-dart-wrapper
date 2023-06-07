package bridge

import (
	"sync"

	librsyncdartwrapper "github.com/filemingo/librsync-dart-wrapper"
	"github.com/filemingo/librsync-dart-wrapper/logging"
	log "github.com/sirupsen/logrus"
)

type Stream interface {
	Send([]byte, int)
}

func ComputeDelta(signatureStr string, targetFilePath string, stream Stream) {
	librsyncdartwrapper.ComputeDelta(signatureStr, targetFilePath, func(b []byte, i int) {
		log.Debugf("[librsync]: Got bytes to send to Java")
		stream.Send(b, i)
	})
}

// For android
type Logger logging.Logger

func AddLogger(logger Logger) {
	logging.AddLogger(logger)
}

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		logging.LogInit()
		log.SetLevel(log.DebugLevel)
	})
}
